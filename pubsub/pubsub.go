package pubsub

import (
	"context"
	"encoding/json"

	gcloud "cloud.google.com/go/pubsub"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/dao/exchange"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/ninjadotorg/constant-api-service/service"
)

const (
	orderTopic = "order"

	orderBookTopic   = "orderbook"
	orderBookSubName = "orderbook-rest-api"
)

type Pubsub struct {
	c           *gcloud.Client
	exchangeDAO *exchange.Exchange
	bc          *blockchain.Blockchain
	logger      *zap.Logger

	// used to hold all ws conns in mem
	subscribers map[*Subscriber]bool
	register    chan *Subscriber
	unregister  chan *Subscriber

	// used to communicate with subscribers
	message chan []byte
}

func New(c *gcloud.Client, exchangeDAO *exchange.Exchange, bc *blockchain.Blockchain, logger *zap.Logger) *Pubsub {
	ps := &Pubsub{
		c:           c,
		exchangeDAO: exchangeDAO,
		bc:          bc,
		logger:      logger,

		subscribers: make(map[*Subscriber]bool),
		register:    make(chan *Subscriber),
		unregister:  make(chan *Subscriber),
		message:     make(chan []byte, 1024),
	}

	go ps.subscribeToOrderBookTopic()
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

func (p *Pubsub) subscribeToOrderBookTopic() {
	t := p.c.Topic(orderBookTopic)
	sub := p.c.Subscription(orderBookSubName)

	exists, err := sub.Exists(context.Background())
	if err != nil {
		p.logger.Error("p.c.CreateSubscription", zap.Error(err))
		return
	}
	if !exists {
		sub, err = p.c.CreateSubscription(context.Background(), orderBookSubName, gcloud.SubscriptionConfig{Topic: t})
		if err != nil {
			p.logger.Error("p.c.CreateSubscription", zap.Error(err))
			return
		}
	}
	err = sub.Receive(context.Background(), func(ctx context.Context, m *gcloud.Message) {
		p.logger.Debug("Receive message", zap.ByteString("data", m.Data))

		if err := p.handleOrderBookMsg(m.Data); err != nil {
			p.logger.Error("p.handleOrderBookMsg", zap.ByteString("data", m.Data), zap.String("error", err.Error()))
			// m.Nack()
			// return
		}

		m.Ack()
	})
	if err != nil {
		p.logger.Error("sub.Receive", zap.Error(err))
	}
}

func (p *Pubsub) handleOrderBookMsg(m []byte) error {
	var body struct {
		Data interface{}
		Type string
	}
	if err := json.Unmarshal(m, &body); err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	j, _ := json.Marshal(body.Data)
	switch body.Type {
	case "change":
		p.message <- j
		return nil
	case "match":
		var msg serializers.OrderBookMatchMsg
		if err := json.Unmarshal(j, &msg); err != nil {
			return errors.Wrap(err, "json.Unmarshal")
		}
		return errors.Wrap(p.handleMatchOrderBook(&msg), "p.handleMatchOrderBook")
	default:
		return errors.Errorf("unsuppoted type %s", body.Type)
	}
}

func (p *Pubsub) handleMatchOrderBook(data *serializers.OrderBookMatchMsg) error {
	takerOrder, err := p.exchangeDAO.FindOrderByID(data.TakerOrderID)
	if err != nil {
		return errors.Wrapf(err, "p.exchangeSvc.FindOrderByID %d", data.TakerOrderID)
	}
	makerOrder, err := p.exchangeDAO.FindOrderByID(data.MakerOrderID)
	if err != nil {
		return errors.Wrapf(err, "p.exchangeSvc.FindOrderByID %d", data.MakerOrderID)
	}

	switch takerOrder.Side {
	case models.Buy:
		err = p.makeTransaction(takerOrder, makerOrder)
	case models.Sell:
		err = p.makeTransaction(makerOrder, takerOrder)
	}
	return errors.Wrap(err, "p.makeTransaction")
}

func (p *Pubsub) makeTransaction(buyer, seller *models.Order) error {
	switch buyer.Market.BaseCurrency.Name {
	case "CONSTANT": // send constant token from buyer to seller
		txID, err := p.bc.Createandsendtransaction(buyer.User.PrivKey, serializers.WalletSend{
			PaymentAddresses: map[string]uint64{
				seller.User.PaymentAddress: buyer.Price,
			},
		})
		if err != nil {
			return errors.Wrap(err, "p.bc.Createandsendtransaction")
		}
		tx, err := service.GetBlockchainTxByHash(txID, 3, p.bc)
		if err != nil {
			return errors.Wrap(err, "p.bc.WaitForTx")
		}
		p.logger.Debug("tx", zap.Any("tx", tx))
	default:
		return errors.Errorf("unsupported currency: %q", buyer.Market.BaseCurrency.Name)
	}

	// send token from seller to buyer
	if err := p.bc.Sendcustomtokentransaction(seller.User.PrivKey, serializers.WalletSend{
		Type:        1,
		TokenID:     buyer.Market.QuoteCurrency.TokenID,
		TokenName:   buyer.Market.QuoteCurrency.TokenName,
		TokenSymbol: buyer.Market.QuoteCurrency.TokenSymbol,
		PaymentAddresses: map[string]uint64{
			buyer.User.PaymentAddress: buyer.Quantity,
		},
	}); err != nil {
		return errors.Wrap(err, "p.bc.Sendcustomtokentransaction")
	}

	if err := p.exchangeDAO.SetFilledOrders(buyer, seller); err != nil {
		return errors.Wrap(err, "p.exchangeDAO.SetFilledOrders")
	}
	return nil
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
