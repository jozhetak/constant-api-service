package service

import (
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/dao/exchange"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
)

type ExchangeService struct {
	r *exchange.Exchange
}

func NewExchange(r *exchange.Exchange) *ExchangeService {
	return &ExchangeService{r}
}

func (e *ExchangeService) ListMarkets(base string) ([]*serializers.MarketResp, error) {
	markets, err := e.r.ListMarkets(base)
	if err != nil {
		return nil, errors.Wrap(err, "c.portalDao.ListByBase")
	}
	return toMarketResp(markets), nil
}

func (e *ExchangeService) CreateOrder(u *models.User, symbol string, price uint64, quantity uint64, typ, side string) (*serializers.OrderResp, error) {
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
		return nil, errors.Wrap(err, "e.portalDao.FindMarketBySymbol")
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
		return nil, errors.Wrap(err, "e.portalDao.CreateOrder")
	}
	return assembleOrder(order), nil
}

func (e *ExchangeService) UserOrderHistory(u *models.User, symbol, status, side string, limit *string, page *string) ([]*serializers.OrderResp, error) {
	if symbol == "" {
		return nil, ErrInvalidSymbol
	}

	var l, p int
	var err error
	if limit != nil && page != nil {
		l, p, err = parsePaginationQuery(*limit, *page)
		if err != nil {
			return nil, errors.Wrap(err, "parsePaginationQuery")
		}
	}

	var oStatus *models.OrderStatus
	if status != "" {
		st := models.GetOrderStatus(status)
		if st == models.InvalidOrderStatus {
			return nil, ErrInvalidOrderStatus
		}
		oStatus = &st
	}

	var oSide *models.OrderSide
	if side != "" {
		si := models.GetOrderSide(side)
		if si == models.InvalidOrderSide {
			return nil, ErrInvalidOrderSide
		}
		oSide = &si
	}

	orders, err := e.r.OrderHistory(symbol, oStatus, oSide, &l, &p, u)
	if err != nil {
		return nil, errors.Wrap(err, "e.portalDao.OrderHistory")
	}
	return toOrderResp(orders), nil
}

func (e *ExchangeService) MarketHistory(symbol, limit, page string) ([]*serializers.OrderResp, error) {
	if symbol == "" {
		return nil, ErrInvalidSymbol
	}
	l, p, err := parsePaginationQuery(limit, page)
	if err != nil {
		return nil, errors.Wrap(err, "parsePaginationQuery")
	}

	status := models.Filled
	orders, err := e.r.OrderHistory(symbol, &status, nil, &l, &p, nil)
	if err != nil {
		return nil, errors.Wrap(err, "e.portalDao.OrderHistory")
	}
	return toOrderResp(orders), nil
}

func (e *ExchangeService) SymbolRates(timeRange string) ([]serializers.SymbolRate, error) {
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
		return nil, errors.Wrap(err, "e.portalDao.SymbolRates")
	}
	return toSymbolRatesResp(rates), nil
}

func (e *ExchangeService) MarketRates() ([]*serializers.MarketRate, error) {
	rates, err := e.r.MarketRates()
	if err != nil {
		return nil, errors.Wrap(err, "e.portalDao.MarketRates")
	}
	return toMarketRatesResp(rates), nil
}

func (e *ExchangeService) FindOrderByID(idS string) (*serializers.OrderResp, error) {
	id, err := strconv.Atoi(idS)
	if err != nil {
		return nil, ErrInvalidOrder
	}
	order, err := e.r.FindOrderByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "e.portalDao.FindByID")
	}
	if order == nil {
		return nil, ErrInvalidOrder
	}
	return assembleOrder(order), nil
}

func toSymbolRatesResp(cs []exchange.SymbolRate) []serializers.SymbolRate {
	resp := make([]serializers.SymbolRate, 0, len(cs))
	for _, cr := range cs {
		resp = append(resp, assembleSymbolRate(cr))
	}
	return resp
}

func toMarketRatesResp(cs []*exchange.MarketRate) []*serializers.MarketRate {
	resp := make([]*serializers.MarketRate, 0, len(cs))
	for _, cr := range cs {
		resp = append(resp, assembleMarketRate(cr))
	}
	return resp
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
		BaseCurrency:         m.BaseCurrency.Name,
		QuoteCurrency:        m.QuoteCurrency.Name,
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

func assembleSymbolRate(m exchange.SymbolRate) serializers.SymbolRate {
	return serializers.SymbolRate{
		SymbolCode: m.SymbolCode,
		Volume:     m.Volume,
		Last:       m.Last,
		High:       m.High,
		Low:        m.Low,
		PrevPrice:  m.PrevPrice,
		PrevVolume: m.PrevVolume,
	}
}

func assembleMarketRate(m *exchange.MarketRate) *serializers.MarketRate {
	return &serializers.MarketRate{
		SymbolCode: m.SymbolCode,
		Last:       m.Last,
		Bid:        m.Bid,
		Ask:        m.Ask,
		Volume:     m.Volume,
		High24h:    m.High24h,
		Low24h:     m.Low24h,
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
