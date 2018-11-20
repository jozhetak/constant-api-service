package service

import (
	"time"

	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/dao/exchange"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
)

type Exchange struct {
	r *exchange.Exchange
}

func NewExchange(r *exchange.Exchange) *Exchange {
	return &Exchange{r}
}

func (e *Exchange) ListMarkets(base string) ([]*serializers.MarketResp, error) {
	markets, err := e.r.ListMarkets(base)
	if err != nil {
		return nil, errors.Wrap(err, "c.r.ListByBase")
	}
	return toMarketResp(markets), nil
}

func (e *Exchange) CreateOrder(u *models.User, symbol string, price float64, quantity uint, typ, side string) (*serializers.OrderResp, error) {
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

func (e *Exchange) UserOrderHistory(u *models.User, symbol, status, limit, page string) ([]*serializers.OrderResp, error) {
	if symbol == "" {
		return nil, ErrInvalidSymbol
	}
	l, p, err := parsePaginationQuery(limit, page)
	if err != nil {
		return nil, errors.Wrap(err, "parsePaginationQuery")
	}

	var oStatus *models.OrderStatus
	if status != "" {
		st := models.GetOrderStatus(status)
		if st == models.InvalidOrderStatus {
			return nil, ErrInvalidOrderStatus
		}
		oStatus = &st
	}

	orders, err := e.r.OrderHistory(symbol, oStatus, l, p, u)
	if err != nil {
		return nil, errors.Wrap(err, "e.r.OrderHistory")
	}
	return toOrderResp(orders), nil
}

func (e *Exchange) MarketHistory(symbol, limit, page string) ([]*serializers.OrderResp, error) {
	if symbol == "" {
		return nil, ErrInvalidSymbol
	}
	l, p, err := parsePaginationQuery(limit, page)
	if err != nil {
		return nil, errors.Wrap(err, "parsePaginationQuery")
	}

	status := models.Filled
	orders, err := e.r.OrderHistory(symbol, &status, l, p, nil)
	if err != nil {
		return nil, errors.Wrap(err, "e.r.OrderHistory")
	}
	return toOrderResp(orders), nil
}

func (e *Exchange) SymbolRates(timeRange string) ([]exchange.SymbolRate, error) {
	var from, to time.Time
	switch timeRange {
	case "1h":
		from, to = time.Now().Add(-1*time.Hour), time.Now()
	case "4h":
		from, to = time.Now().Add(-4*time.Hour), time.Now()
	case "24h":
		from, to = time.Now().Add(-24*time.Hour), time.Now()
	default:
		return nil, ErrInvalidArgument
	}

	rates, err := e.r.SymbolRates(from, to)
	if err != nil {
		return nil, errors.Wrap(err, "e.r.SymbolRates")
	}
	return rates, nil
}

func toMarketResp(cs []*models.Market) []*serializers.MarketResp {
	resp := make([]*serializers.MarketResp, 0, len(cs))
	for _, cr := range cs {
		resp = append(resp, assembleMarket(cr))
	}
	return resp
}

func toOrderResp(cs []*models.Order) []*serializers.OrderResp {
	resp := make([]*serializers.OrderResp, 0, len(cs))
	for _, cr := range cs {
		resp = append(resp, assembleOrder(cr))
	}
	return resp
}

func assembleMarket(m *models.Market) *serializers.MarketResp {
	return &serializers.MarketResp{
		BaseCurrency:         m.BaseCurrency,
		QuoteCurrency:        m.QuoteCurrency,
		DisplayName:          m.DisplayName,
		State:                m.State.String(),
		SymbolCode:           m.SymbolCode,
		Icon:                 m.Icon,
		TradeEnabled:         m.TradeEnabled,
		FeePrecision:         m.FeePrecision,
		TradePricePrecision:  m.TradePricePrecision,
		TradeTotalPrecision:  m.TradeTotalPrecision,
		TradeAmountPrecision: m.TradeAmountPrecision,
	}
}

func assembleOrder(o *models.Order) *serializers.OrderResp {
	return &serializers.OrderResp{
		ID:         o.ID,
		SymbolCode: o.Market.SymbolCode,
		Price:      o.Price,
		Quantity:   o.Quantity,
		Type:       o.Type.String(),
		Status:     o.Status.String(),
		Side:       o.Side.String(),
		Time:       o.Time.Format(time.RFC3339),
	}
}
