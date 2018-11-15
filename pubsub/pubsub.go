package pubsub

import (
	"encoding/json"

	"go.uber.org/zap"
)

type Service struct {
	subscribers map[*Subscriber]bool
	message     chan []byte

	register   chan *Subscriber
	unregister chan *Subscriber

	logger *zap.Logger
}

func NewService() *Service {
	s := &Service{
		subscribers: make(map[*Subscriber]bool),
		register:    make(chan *Subscriber),
		unregister:  make(chan *Subscriber),
		message:     make(chan []byte),
	}
	go s.run()
	return s
}

func (s *Service) Publish(m interface{}) {
	b, err := json.Marshal(m)
	if err != nil {
		s.logger.Error("json.Marshal", zap.Error(err))
		return
	}
	select {
	case s.message <- b:
	default:
	}
}

func (s *Service) run() {
	for {
		select {
		case sub := <-s.register:
			s.subscribers[sub] = true
		case sub := <-s.unregister:
			if _, ok := s.subscribers[sub]; ok {
				close(sub.message)
				delete(s.subscribers, sub)
			}
		case m := <-s.message:
			for sub := range s.subscribers {
				select {
				case sub.message <- m:
				default:
					close(sub.message)
					delete(s.subscribers, sub)
				}
			}
		}
	}
}

type Subscriber struct {
	service *Service
	message chan []byte
}

func NewSubscriber(s *Service) *Subscriber {
	sub := &Subscriber{
		service: s,
		message: make(chan []byte, 1024),
	}
	sub.service.register <- sub
	return sub
}

func (s *Subscriber) Read() <-chan []byte {
	return s.message
}
