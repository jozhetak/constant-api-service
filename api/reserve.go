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
	list = append(list, models.ReserseParty{
		ID:       models.ReservePartyEth,
		Name:     "Ethereum",
		IsActive: true,
	})
	c.JSON(http.StatusOK, serializers.Resp{Result: list})
}

// Create a request reserve with related party
func (server *Server) RequestReserve(c *gin.Context) {

}

// Get detail data of request reserve
func (server *Server) GetRequestReserve(c *gin.Context) {

}

// Get history data of request reserve
func (server *Server) RequestReserveHistory(c *gin.Context) {

}

// Create a return reserve request with related party
func (server *Server) ReturnRequestReserve(c *gin.Context) {

}

// Get detail data of return request reserve
func (server *Server) GetReturnRequestReserve(c *gin.Context) {

}

// Get history data of return request reserve
func (server *Server) ReturnRequestReserveHistory(c *gin.Context) {

}
