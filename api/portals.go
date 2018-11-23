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
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}

	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	b, err := s.portalSvc.CreateBorrow(user, req)
	if err != nil {
		s.logger.Error("s.borrowSvc.Create", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: b})
}

func (s *Server) ListBorrowsByUser(c *gin.Context) {
	var (
		paymentAddress = c.DefaultQuery("payment_address", "")
		state          = c.DefaultQuery("state", "")
		page           = c.DefaultQuery("page", "1")
		limit          = c.DefaultQuery("limit", "10")
	)

	bs, err := s.portalSvc.ListBorrowsByUser(paymentAddress, state, limit, page)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidBorrowState, service.ErrInvalidLimit, service.ErrInvalidPage:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: bs})
	default:
		s.logger.Error("s.borrowSvc.ListByUser", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (s *Server) ListAllBorrows(c *gin.Context) {
	var (
		state = c.DefaultQuery("state", "")
		page  = c.DefaultQuery("page", "1")
		limit = c.DefaultQuery("limit", "10")
	)

	bs, err := s.portalSvc.ListAllBorrows(state, limit, page)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidBorrowState, service.ErrInvalidLimit, service.ErrInvalidPage:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: bs})
	default:
		s.logger.Error("s.borrowSvc", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (s *Server) FindByID(c *gin.Context) {
	b, err := s.portalSvc.FindBorrowByID(c.Param("id"))
	switch cErr := errors.Cause(err); cErr {
	case service.ErrBorrowNotFound:
		c.JSON(http.StatusNotFound, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: service.AssembleBorrow(b)})
	default:
		s.logger.Error("s.borrowSvc.FindByID", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})

	}
}

func (s *Server) ProcessStateBorrowByID(c *gin.Context) {
	b, err := s.portalSvc.FindBorrowByID(c.Param("id"))
	switch cErr := errors.Cause(err); cErr {
	case service.ErrBorrowNotFound:
		c.JSON(http.StatusNotFound, serializers.Resp{Error: cErr.(*service.Error)})
	}

	result, err := s.portalSvc.UpdateStatusBorrowRequest(b, c.DefaultQuery("action", ""), c.DefaultQuery("costant_loan_response_tx_id", ""))
	switch cErr := errors.Cause(err); cErr {
	case service.ErrBorrowNotFound:
		c.JSON(http.StatusNotFound, serializers.Resp{Error: cErr.(*service.Error)})
	default:
		c.JSON(http.StatusOK, serializers.Resp{Result: result})
	}
}

func (s *Server) PayBorrowByID(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	// call blockchain to check tx payment
	b, err := s.portalSvc.FindBorrowByID(c.Param("id"))
	switch cErr := errors.Cause(err); cErr {
	case service.ErrBorrowNotFound:
		c.JSON(http.StatusNotFound, serializers.Resp{Error: cErr.(*service.Error)})
	}
	paymentTx, err := s.portalSvc.PaymentTxForLoanRequest(user, b, c.DefaultQuery("costant_payment_tx_id", ""))
	if err != nil {
		c.JSON(http.StatusBadGateway, serializers.Resp{Error: err})
		return
	}
	if paymentTx != nil {
		switch b.CollateralType {
		case "ETH":
			// TODO call web3 to process collateral
			//
			//

		}
		c.JSON(http.StatusOK, serializers.Resp{Result: true})
	} else {
		c.JSON(http.StatusOK, serializers.Resp{Result: false})
	}
}

func (s *Server) WithdrawBorrowByID(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	// call blockchain to check tx payment
	b, err := s.portalSvc.FindBorrowByID(c.Param("id"))
	switch cErr := errors.Cause(err); cErr {
	case service.ErrBorrowNotFound:
		c.JSON(http.StatusNotFound, serializers.Resp{Error: cErr.(*service.Error)})
		return
	}
	key := c.DefaultQuery("key", "")
	if len(key) == 0 {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}
	// process with constant network
	withdrawTx, err := s.portalSvc.WithdrawTxForLoanRequest(user, b, key)
	if err != nil {
		c.JSON(http.StatusBadGateway, serializers.Resp{Error: err})
		return
	}
	if withdrawTx != nil {
		switch b.CollateralType {
		case "ETH":
			// TODO call web3 to process collateral
			//
			//

		}
		c.JSON(http.StatusOK, serializers.Resp{Result: true})
	} else {
		c.JSON(http.StatusOK, serializers.Resp{Result: false})
	}
}

func (s *Server) GetLoanParams(c *gin.Context) {
	result, err := s.portalSvc.GetLoanParams()
	c.JSON(http.StatusOK, serializers.Resp{Error: err, Result: result})
}
