package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Get list Party which support for reserve service
func (s *server) GetReserveParty(c *gin.Context) {
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
func (s *server) CreateContribution(c *gin.Context) {
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

	contribution, err := s.reserveSvc.CreateContribution(&user, &req)

	if err {
		c.JSON(http.StatusOK, serializers.Resp{Error: err})
	}

	return c.JSON(http.StatusOK, serializers.Resp{Result: contribution, Error: nil})
}

// Get detail data of request reserve
func (s *server) GetContribution(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	contribution, err := s.reserveSvc.GetContributionById(&user, c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusOK, serializers.Resp{Error: err})
	}
	return contribution, nil
}

// Get history data of request reserve
func (s *server) ContributionHistory(c *gin.Context) {
	var (
		page  = c.DefaultQuery("page", "1")
		limit = c.DefaultQuery("limit", "10")
	)

	var filter map[string]interface{}

	contributions, err := s.reserveSvc.GetContributions(state, &filter, limit, page)
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
func (s *server) CreateDisbursement(c *gin.Context) {
	// TODO
	s.reserveSvc.CreateDisbursement(&serializers.ReserveDisbursementRequest{})
}

// Get detail data of return request reserve
func (s *server) GetDisbursement(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	disbursement, err := s.reserveSvc.GetDisbursementById(&user, c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusOK, serializers.Resp{Error: err})
	}
	return disbursement, nil
}

// Get history data of return request reserve
func (s *server) DisbursementHistory(c *gin.Context) {
	var (
		page  = c.DefaultQuery("page", "1")
		limit = c.DefaultQuery("limit", "10")
	)

	var filter map[string]interface{}

	disbursements, err := s.reserveSvc.GetDisbursements(state, &filter, limit, page)
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

func (s *server) PrimetrustWebHook(c *gin.Context) {
	// TODO
	// get request and update data for payment party
	// call blockchain network to for contribution case
	s.reserveSvc.PrimetrustWebHook(nil)
}
