package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/pubsub"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
)

func (s *Server) ListMarkets(c *gin.Context) {
	markets, err := s.exchangeSvc.ListMarkets(c.Query("base"))
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

	order, err := s.exchangeSvc.CreateOrder(user, req.Symbol, req.Price, req.Quantity, req.Type, req.Side)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidOrderSide, service.ErrInvalidOrderType, service.ErrInvalidSymbol:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		go s.ps.Publish(order)
		c.JSON(http.StatusOK, serializers.Resp{Result: order})
	default:
		s.logger.Error("s.exchangeSvc.CreateOrder", zap.Error(err))
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
