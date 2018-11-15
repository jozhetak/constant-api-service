package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/constant-api-service/service"
	"go.uber.org/zap"
)

func (s *Server) ListAccounts(c *gin.Context) {
	v, err := s.walletSvc.ListAccounts(c.DefaultQuery("params", ""))
	if err != nil {
		s.logger.Error("s.walletSvc.ListAccounts", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, v)
}

func (s *Server) GetBalanceByPrivateKey(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
		return
	}

	v, err := s.walletSvc.GetBalanceByPrivateKey(user.PrivKey)
	if err != nil {
		s.logger.Error("s.walletSvc.ListAccounts", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, v)
}
