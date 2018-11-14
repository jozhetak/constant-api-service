package api

import (
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	// maxMessageSize = 512
)

type hub struct {
	clients map[*client]bool

	register   chan *client
	unregister chan *client

	message chan []byte
}

func newHub() *hub {
	return &hub{
		clients:    make(map[*client]bool),
		register:   make(chan *client),
		unregister: make(chan *client),
		message:    make(chan []byte),
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.send)
			}
		case m := <-h.message:
			for c := range h.clients {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(h.clients, c)
				}
			}
		}
	}
}

type client struct {
	hub    *hub
	logger *zap.Logger
	conn   *websocket.Conn
	send   chan []byte
}

func (c *client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := c.conn.Close(); err != nil {
			c.logger.Error("c.conn.Close", zap.Error(err))
		}
	}()

	for {
		select {
		case b, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(b)
			if err != nil {
				return
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
