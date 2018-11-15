package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
)

func (s *Server) CreateNewBorrow(c *gin.Context) {
	var req serializers.BorrowReq
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

	b, err := s.portalSvc.CreateBorrow(user, req.Amount, req.Hash, req.TxID, req.PaymentAddress)
	if err != nil {
		s.logger.Error("s.borrowSvc.Create", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
		return
	}

	c.JSON(http.StatusOK, response{Data: b})
}

func (s *Server) ListBorrowsByUser(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
		return
	}

	var (
		state = c.DefaultQuery("state", "")
		page  = c.DefaultQuery("page", "1")
		limit = c.DefaultQuery("limit", "10")
	)

	bs, err := s.portalSvc.ListBorrowsByUser(user, state, limit, page)
	switch err := errors.Cause(err); err {
	case service.ErrInvalidBorrowState, service.ErrInvalidLimit, service.ErrInvalidPage:
		c.JSON(http.StatusBadRequest, response{Error: err.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, response{Data: bs})
	default:
		s.logger.Error("s.borrowSvc.ListByUser", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
	}
}

func (s *Server) ListAllBorrows(c *gin.Context) {
	var (
		state = c.DefaultQuery("state", "")
		page  = c.DefaultQuery("page", "1")
		limit = c.DefaultQuery("limit", "10")
	)

	bs, err := s.portalSvc.ListAllBorrows(state, limit, page)
	switch err := errors.Cause(err); err {
	case service.ErrInvalidBorrowState, service.ErrInvalidLimit, service.ErrInvalidPage:
		c.JSON(http.StatusBadRequest, response{Error: err.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, response{Data: bs})
	default:
		s.logger.Error("s.borrowSvc", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})
	}
}

func (s *Server) FindByID(c *gin.Context) {
	borrow, err := s.portalSvc.FindBorrowByID(c.Param("id"))
	switch err := errors.Cause(err); err {
	case service.ErrBorrowNotFound:
		c.JSON(http.StatusNotFound, response{Error: err.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, response{Data: borrow})
	default:
		s.logger.Error("s.borrowSvc.FindByID", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response{Error: service.ErrInternalServerError})

	}
}
