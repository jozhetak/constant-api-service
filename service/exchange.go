package service

import (
	"time"

	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/dao"
	"github.com/ninjadotorg/constant-api-service/models"
)

type Exchange struct {
	r *dao.Exchange
}

func NewExchange(r *dao.Exchange) *Exchange {
	return &Exchange{r}
}

func (e *Exchange) ListMarkets(base string) ([]*Market, error) {
	markets, err := e.r.ListMarkets(base)
	if err != nil {
		return nil, errors.Wrap(err, "c.r.ListByBase")
	}
	return e.transformToResp(markets), nil
}

func (e *Exchange) CreateOrder(u *models.User, symbol string, price float64, quantity uint, typ, side string) (*Order, error) {
	oTyp := models.GetOrderType(typ)
	if oTyp == models.InvalidOrderType {
		return nil, ErrInvalidOrderType
	}

	oSide := models.GetOrderSide(side)
	if oSide == models.InvalidOrderSide {
		return nil, ErrInvalidOrderSide
	}

	market, err := e.r.FindMarketBySymbol(symbol)
	if err != nil {
		return nil, errors.Wrap(err, "e.r.FindMarketBySymbol")
	}
	if market == nil {
		return nil, ErrInvalidSymbol
	}

	order, err := e.r.CreateOrder(&models.Order{
		User:     u,
		Market:   market,
		Price:    price,
		Quantity: quantity,
		Type:     oTyp,
		Status:   models.New,
		Side:     oSide,
		Time:     time.Now(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "e.r.CreateOrder")
	}
	return assembleOrder(order), nil
}

func (e *Exchange) transformToResp(cs []*models.Market) []*Market {
	resp := make([]*Market, 0, len(cs))
	for _, cr := range cs {
		resp = append(resp, assembleMarket(cr))
	}
	return resp
}

type Market struct {
	BaseCurrency   string `json:"base_currency"`
	MarketCurrency string `json:"market_currency"`
	Symbol         string `json:"symbol"`
}

func assembleMarket(c *models.Market) *Market {
	return &Market{
		BaseCurrency:   c.BaseCurrency,
		Symbol:         c.Symbol,
		MarketCurrency: c.MarketCurrency,
	}
}

type Order struct {
	ID       uint    `json:"id"`
	Symbol   string  `json:"symbol"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
	Type     string  `json:"type"`
	Status   string  `json:"status"`
	Side     string  `json:"side"`
	Time     string  `json:"time"`
}

func assembleOrder(o *models.Order) *Order {
	return &Order{
		ID:       o.ID,
		Symbol:   o.Market.Symbol,
		Price:    o.Price,
		Quantity: o.Quantity,
		Type:     o.Type.String(),
		Status:   o.Status.String(),
		Side:     o.Side.String(),
		Time:     o.Time.Format(time.RFC3339),
	}
}
