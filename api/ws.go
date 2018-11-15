package api

import (
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/pubsub"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type ws struct {
	subscriber *pubsub.Subscriber
	conn       *websocket.Conn
	logger     *zap.Logger
}

func newWS(sub *pubsub.Subscriber, conn *websocket.Conn, logger *zap.Logger) *ws {
	return &ws{
		subscriber: sub,
		conn:       conn,
		logger:     logger,
	}
}

func (w *ws) sendMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := w.conn.Close(); err != nil {
			w.logger.Error("c.conn.Close", zap.Error(err))
		}
	}()

	for {
		select {
		case b, ok := <-w.subscriber.Read():
			_ = w.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				if err := w.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					w.logger.Error("w.conn.WriteMessage", zap.Error(err))
				}
				return
			}

			w, err := w.conn.NextWriter(websocket.TextMessage)
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
			_ = w.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := w.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
