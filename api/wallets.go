package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
)

func (s *Server) ListAccounts(c *gin.Context) {
	v, err := s.walletSvc.ListAccounts(c.DefaultQuery("params", ""))
	if err != nil {
		s.logger.Error("s.walletSvc.ListAccounts", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, v)
}

func (s *Server) GetCoinBalance(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	v, err := s.walletSvc.GetBalanceByPrivateKey(user.PrivKey)
	if err != nil {
		s.logger.Error("s.walletSvc.ListAccounts", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: v})
}

func (s *Server) GetCoinAndCustomTokenBalance(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	privKey := user.PrivKey
	paymentAddress := user.PaymentAddress
	resp, err := s.walletSvc.GetCoinAndCustomTokenBalance(privKey, paymentAddress)
	if err != nil {
		s.logger.Error("s.walletSvc.ListAccounts", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: resp})
}
