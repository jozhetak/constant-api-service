package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Get list Party which support for reserve service
func (s *Server) GetReserveParty(c *gin.Context) {
	list := []models.ReserseParty{}
	list = append(list, models.ReserseParty{
		ID:       models.ReservePartyPrimeTrust,
		Name:     "Prime Trust",
		IsActive: true,
	})
	/*list = append(list, models.ReserseParty{
		ID:       models.ReservePartyEth,
		Name:     "Ethereum",
		IsActive: true,
	})*/
	c.JSON(http.StatusOK, serializers.Resp{Result: list})
}

// Create a request reserve with related party
func (s *Server) CreateContribution(c *gin.Context) {
	// TODO
	// Validate request
	// TODO

	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	var req serializers.ReserveContributionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}

	contribution, err := s.reserveSvc.CreateContribution(user, &req)

	if err != nil {
		c.JSON(http.StatusOK, serializers.Resp{Error: err})
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: contribution, Error: nil})
}

// Get detail data of request reserve
func (s *Server) GetContribution(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	id := c.Param("id")
	idi, _ := strconv.Atoi(id)

	contribution, err := s.reserveSvc.GetContribution(user, idi)

	if err != nil {
		c.JSON(http.StatusOK, serializers.Resp{Error: err})
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: contribution, Error: nil})
}

// Get history data of request reserve
func (s *Server) ContributionHistory(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	var (
		page  = c.DefaultQuery("page", "1")
		limit = c.DefaultQuery("limit", "10")
	)

	pageI, _ := strconv.Atoi(page)
	limitI, _ := strconv.Atoi(limit)

	var filter map[string]interface{}

	contributions, err := s.reserveSvc.GetContributions(user, &filter, pageI, limitI)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidBorrowState, service.ErrInvalidLimit, service.ErrInvalidPage:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: contributions})
	default:
		s.logger.Error("s.reserveSvc", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

// Create a return reserve request with related party
func (s *Server) CreateDisbursement(c *gin.Context) {
	// TODO
	s.reserveSvc.CreateDisbursement(&serializers.ReserveDisbursementRequest{})
}

// Get detail data of return request reserve
func (s *Server) GetDisbursement(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	id := c.Param("id")
	idi, _ := strconv.Atoi(id)

	disbursement, err := s.reserveSvc.GetDisbursement(user, idi)

	if err != nil {
		c.JSON(http.StatusOK, serializers.Resp{Error: err})
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: disbursement, Error: nil})
}

// Get history data of return request reserve
func (s *Server) DisbursementHistory(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	var (
		page  = c.DefaultQuery("page", "1")
		limit = c.DefaultQuery("limit", "10")
	)
	var filter map[string]interface{}

	pageI, _ := strconv.Atoi(page)
	limitI, _ := strconv.Atoi(limit)

	disbursements, err := s.reserveSvc.GetDisbursements(user, &filter, limitI, pageI)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidBorrowState, service.ErrInvalidLimit, service.ErrInvalidPage:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: disbursements})
	default:
		s.logger.Error("s.reserveSvc", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (s *Server) PrimetrustWebHook(c *gin.Context) {
	// TODO
	// get request and update data for payment party
	// call blockchain network to for contribution case
	s.reserveSvc.PrimetrustWebHook(nil)
}
