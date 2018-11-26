package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
	"net/http"
	"github.com/ninjadotorg/constant-api-service/models"
)

func (server *Server) RegisterBoardCandidate(c *gin.Context) {
	user, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	req := serializers.VotingBoardCandidateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}
	if req.PaymentAddress == "" {
		req.PaymentAddress = user.PaymentAddress
	}
	votingBoardCandidate, err := server.votingSvc.RegisterBoardCandidate(user, models.BoardCandidateType(req.BoardType), req.PaymentAddress)
	if err != nil {
		server.logger.Error("s.votingSvc.RegisterBoardCandidate", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	result := serializers.NewVotingBoardCandidateResp(votingBoardCandidate)
	c.JSON(http.StatusOK, serializers.Resp{Result: result})
}
