package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
	"net/http"
	"github.com/ninjadotorg/constant-api-service/models"
	"strconv"
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

func (server *Server) GetCandidatesList(c *gin.Context) {
	_, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	boardQuery := c.DefaultQuery("board", "0")
	board, _ := strconv.Atoi(boardQuery)
	list, err := server.votingSvc.GetCandidatesList(board, c.DefaultQuery("board", "payment_address"))
	if err != nil {
		server.logger.Error("s.votingSvc.GetCandidatesList", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	result := serializers.VotingBoardCandidateRespList{}
	if len(list) > 0 {
		result = *(serializers.NewVotingBoardCandidateListResp(list))
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: result})
}

func (server *Server) VoteCandidateBoard(c *gin.Context) {
	// TODO

	err := server.votingSvc.VoteCandidateBoard()
	if err != nil {
		server.logger.Error("s.voting.VoteCandidateBoard", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
}

func (server *Server) CreateProposal(c *gin.Context) {
	// TODO
	user, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	req := serializers.VotingProposalRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}
	proposal, err := server.votingSvc.CreateProposal(user, &req)
	if err != nil {
		server.logger.Error("s.votingSvc.CreateProposal", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	switch proposal.GetType() {
	case 1:
		{
			result := serializers.NewProposalDCBResp(proposal.(models.VotingProposalDCB))
			c.JSON(http.StatusOK, serializers.Resp{Result: result})
		}
	case 2:
		{
			result := serializers.NewProposalGOVResp(proposal.(models.VotingProposalGOV))
			c.JSON(http.StatusOK, serializers.Resp{Result: result})
		}
	default:
		{
			server.logger.Error("s.votingSvc.CreateProposal", zap.Error(err))
			c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
			return
		}
	}
}

func (server *Server) GetProposalsList(c *gin.Context) {
	// TODO
	server.votingSvc.GetProposalsList()
}

func (server *Server) GetProposal(c *gin.Context) {
	// TODO
	server.votingSvc.GetProposal()
}

func (server *Server) VoteProposal(c *gin.Context) {
	// TODO
	err := server.votingSvc.VoteProposal()
	if err != nil {
		server.logger.Error("s.voting.VoteProposal", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
}
