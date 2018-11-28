package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/pubsub"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
)

func (s *Server) ListMarkets(c *gin.Context) {
	markets, err := s.exchangeSvc.ListMarkets(nil, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: markets})
}

func (s *Server) CreateOrder(c *gin.Context) {
	var req serializers.OrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}

	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	order, err := s.exchangeSvc.CreateOrder(user, &req)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidOrderSide, service.ErrInvalidOrderType, service.ErrInvalidSymbol:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		go s.ps.Publish(toOrderMsg("add", order))
		c.JSON(http.StatusOK, serializers.Resp{Result: order})
	default:
		s.logger.Error("s.exchangeSvc.CreateOrder", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (s *Server) UserOrderHistory(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	var (
		symbol = c.Query("symbol_code")
		status = c.DefaultQuery("status", "new")
		side   = c.Query("side")
		page   = c.DefaultQuery("page", "1")
		limit  = c.DefaultQuery("limit", "10")
	)
	orders, err := s.exchangeSvc.UserOrderHistory(user, symbol, status, side, &limit, &page)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidSymbol, service.ErrInvalidOrderStatus, service.ErrInvalidLimit, service.ErrInvalidPage:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: orders})
	default:
		s.logger.Error("s.exchangeSvc.OrderHistory", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (s *Server) MarketHistory(c *gin.Context) {
	var (
		symbol = c.Query("symbol_code")
		page   = c.DefaultQuery("page", "1")
		limit  = c.DefaultQuery("limit", "10")
	)

	orders, err := s.exchangeSvc.MarketHistory(symbol, limit, page)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidSymbol, service.ErrInvalidLimit, service.ErrInvalidPage:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: orders})
	default:
		s.logger.Error("s.exchangeSvc.OrderHistory", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (s *Server) SymbolRates(c *gin.Context) {
	rates, err := s.exchangeSvc.SymbolRates(c.DefaultQuery("range", "24h"))
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidArgument:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: rates})
	default:
		s.logger.Error("s.exchangeSvc.SymbolRates", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (s *Server) MarketRates(c *gin.Context) {
	rates, err := s.exchangeSvc.MarketRates()
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidArgument:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: rates})
	default:
		s.logger.Error("s.exchangeSvc.SymbolRates", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (s *Server) ExchangeWS(c *gin.Context) {
	conn, err := s.up.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		s.logger.Error("s.up.Upgrade", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	logger := s.logger.With(zap.String("module", "client"))
	ws := newWS(pubsub.NewSubscriber(s.ps), conn, logger)
	go ws.sendMessage()
}

func toOrderMsg(typ string, o *serializers.OrderResp) *serializers.OrderPubMsg {
	ts, _ := time.Parse(time.RFC3339, o.Time)
	return &serializers.OrderPubMsg{
		Type: typ,
		Order: &serializers.OrderMsg{
			ID:     int(o.ID),
			Price:  o.Price,
			Size:   o.Quantity,
			Side:   strings.ToLower(o.Side),
			Symbol: o.SymbolCode,
			Type:   o.Type,
			Time:   ts.Unix(),
		},
	}
}
