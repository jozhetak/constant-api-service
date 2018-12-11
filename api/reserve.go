package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
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
	/*  list = append(list, models.ReserseParty{
		ID:       models.ReservePartyEth,
		Name:     "Ethereum",
		IsActive: true,
	})*/
	c.JSON(http.StatusOK, serializers.Resp{Result: list})
}

// Create a request reserve with related party
func (s *Server) CreateContribution(c *gin.Context) {
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
		fmt.Println("create contribution fail", err)
		c.JSON(http.StatusOK, serializers.Resp{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: contribution, Error: nil})
}

// Get detail data of request reserve
func (s *Server) GetContribution(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	contribution, err := s.reserveSvc.GetContribution(id)
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

	page, limit := s.pagingFromContext(c)

	filter := map[string]interface{}{
		"user_id": user.ID,
	}

	contributions, err := s.reserveSvc.GetContributions(&filter, page, limit)
	if err != nil {
		s.logger.Error("s.reserveSvc", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: contributions})
}

// Create a return reserve request with related party
func (s *Server) CreateDisbursement(c *gin.Context) {
	user, err := s.userFromContext(c)
	if err != nil {
		s.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	var req serializers.ReserveDisbursementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}

	disbursement, err := s.reserveSvc.CreateDisbursement(user, &req)

	if err != nil {
		c.JSON(http.StatusOK, serializers.Resp{Error: err})
		return
	}

	c.JSON(http.StatusOK, serializers.Resp{Result: disbursement, Error: nil})
}

// Get detail data of return request reserve
func (s *Server) GetDisbursement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	disbursement, err := s.reserveSvc.GetDisbursement(id)
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

	page, limit := s.pagingFromContext(c)

	filter := map[string]interface{}{
		"user_id": user.ID,
	}
	disbursements, err := s.reserveSvc.GetDisbursements(&filter, limit, page)
	if err != nil {
		s.logger.Error("s.reserveSvc", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: disbursements})
}

func (s *Server) PrimetrustWebHook(c *gin.Context) {
	var req serializers.PrimetrustChangedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}
	err := s.reserveSvc.PrimetrustWebHook(&req)
	if err != nil {
		fmt.Println("s.reserveSvc.PrimetrustWebHook err", err.Error())
	}
	return
}
