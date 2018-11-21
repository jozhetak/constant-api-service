package pubsub

import (
	"context"
	"encoding/json"

	gcloud "cloud.google.com/go/pubsub"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/serializers"
)

const (
	orderTopic     = "order"
	orderBookTopic = "orderbook"
)

type Pubsub struct {
	c      *gcloud.Client
	logger *zap.Logger

	subscribers map[*Subscriber]bool
	register    chan *Subscriber
	unregister  chan *Subscriber
	message     chan []byte
}

func New(c *gcloud.Client, logger *zap.Logger) *Pubsub {
	ps := &Pubsub{
		c:      c,
		logger: logger,

		subscribers: make(map[*Subscriber]bool),
		register:    make(chan *Subscriber),
		unregister:  make(chan *Subscriber),
		message:     make(chan []byte, 1024),
	}

	go ps.subscribe()
	go ps.handleSubscribers()

	return ps
}

func (p *Pubsub) pubOrder(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		p.logger.Error("json.Marshal", zap.Error(err))
		return
	}
	_, err = p.c.Topic(orderTopic).Publish(context.Background(), &gcloud.Message{
		Data: b,
	}).Get(context.Background())

	if err != nil {
		p.logger.Error("p.c.Topic.Publish", zap.String("topic", orderTopic), zap.Error(err))
	}
}

func (p *Pubsub) Publish(v interface{}) {
	switch v.(type) {
	case *serializers.OrderPubMsg:
		p.logger.Debug("publish order", zap.Any("order", v))
		p.pubOrder(v)
	default:
		p.logger.Warn("unsupport v type", zap.Any("v", v))
	}
}

func (p *Pubsub) subscribe() {
	p.subscribeToTopic(p.c.Topic(orderBookTopic))
}

func (p *Pubsub) subscribeToTopic(t *gcloud.Topic) {
	sub := p.c.Subscription("restapiorder")
	exists, err := sub.Exists(context.Background())
	if err != nil {
		p.logger.Error("p.c.CreateSubscription", zap.Error(err))
		return
	}
	if !exists {
		p.logger.Debug("create sub")
		sub, err = p.c.CreateSubscription(context.Background(), "restapiorder", gcloud.SubscriptionConfig{Topic: t})
		if err != nil {
			p.logger.Error("p.c.CreateSubscription", zap.Error(err))
			return
		}
	}
	err = sub.Receive(context.Background(), func(ctx context.Context, m *gcloud.Message) {
		p.logger.Debug("Receive message", zap.ByteString("data", m.Data))

		p.message <- m.Data
		m.Ack()
	})
	if err != nil {
		p.logger.Error("sub.Receive", zap.Error(err))
	}
}

func (p *Pubsub) handleSubscribers() {
	for {
		select {
		case sub := <-p.register:
			p.subscribers[sub] = true
		case sub := <-p.unregister:
			if _, ok := p.subscribers[sub]; ok {
				close(sub.message)
				delete(p.subscribers, sub)
			}
		case m := <-p.message:
			for sub := range p.subscribers {
				select {
				case sub.message <- m:
				default:
					close(sub.message)
					delete(p.subscribers, sub)
				}
			}
		}
	}
}

type Subscriber struct {
	ps      *Pubsub
	message chan []byte
}

func NewSubscriber(ps *Pubsub) *Subscriber {
	sub := &Subscriber{
		ps:      ps,
		message: make(chan []byte, 1024),
	}
	sub.ps.register <- sub
	return sub
}

func (s *Subscriber) Read() <-chan []byte {
	return s.message
}
