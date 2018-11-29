package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"net/http"
)

// Get list Party which support for reserve service
func (server *Server) GetReserveParty(c *gin.Context) {
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
func (server *Server) CreateContribution(c *gin.Context) {

	// TODO
	// Validate request
	server.reserveSvc.CreateContribution(&serializers.ReserveContributionRequest{})
	// TODO
}

// Get detail data of request reserve
func (server *Server) GetContribution(c *gin.Context) {
	// TODO
	server.reserveSvc.GetContribution(nil)
}

// Get history data of request reserve
func (server *Server) ContributionHistory(c *gin.Context) {
	// TODO
	server.reserveSvc.ContributionHistory(nil)
}

// Create a return reserve request with related party
func (server *Server) CreateDisbursement(c *gin.Context) {
	// TODO
	server.reserveSvc.CreateDisbursement(&serializers.ReserveDisbursementRequest{})
}

// Get detail data of return request reserve
func (server *Server) GetDisbursement(c *gin.Context) {
	// TODO
	server.reserveSvc.GetDisbursement(nil)
}

// Get history data of return request reserve
func (server *Server) DisbursementHistory(c *gin.Context) {
	// TODO
	server.reserveSvc.DisbursementHistory(nil)
}

func (server *Server) PrimetrustWebHook(c *gin.Context) {
	// TODO
	// get request and update data for payment party
	// call blockchain network to for contribution case
	server.reserveSvc.PrimetrustWebHook(nil)
}
