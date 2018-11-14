package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/service"
)

func (s *Server) ListMarkets(c *gin.Context) {
	currencies, err := s.exchangeSvc.ListMarkets(c.Query("base"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, response{Data: currencies})
}

func (s *Server) CreateOrder(c *gin.Context) {
	type request struct {
		Symbol   string  `json:"symbol" binding:"required"`
		Price    float64 `json:"price" binding:"required"`
		Quantity uint    `json:"quantity" binding:"required"`
		Type     string  `json:"type" binding:"required"`
		Side     string  `json:"side" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response{Error: service.ErrInvalidArgument})
		return
	}

	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
		return
	}

	order, err := s.exchangeSvc.CreateOrder(user, req.Symbol, req.Price, req.Quantity, req.Type, req.Side)
	switch err := errors.Cause(err); err {
	case service.ErrInvalidOrderSide, service.ErrInvalidOrderType, service.ErrInvalidSymbol:
		c.JSON(http.StatusBadRequest, response{Error: err.(*service.Error)})
	case nil:
		go func() {
			b, err := json.Marshal(order)
			if err != nil {
				s.logger.Error("json.Marshal", zap.Error(err))
				return
			}
			select {
			case s.hub.message <- b:
			default:
			}
		}()
		c.JSON(http.StatusOK, response{Data: order})
	default:
		s.logger.Error("s.exchangeSvc.CreateOrder", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
	}
}

func (s *Server) ExchangeWS(c *gin.Context) {
	conn, err := s.up.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		s.logger.Error("s.up.Upgrade", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
		return
	}

	client := &client{
		hub:    s.hub,
		logger: s.logger.With(zap.String("module", "client")),
		conn:   conn,
		send:   make(chan []byte, 1024),
	}
	client.hub.register <- client

	go client.write()
}
