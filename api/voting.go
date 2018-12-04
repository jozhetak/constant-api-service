package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service"
)

func (server *Server) RegisterBoardCandidate(c *gin.Context) {
	user, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	req := serializers.RegisterBoardCandidateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}
	if req.PaymentAddress == "" {
		req.PaymentAddress = user.PaymentAddress
	}
	votingBoardCandidate, err := server.votingSvc.RegisterBoardCandidate(user, models.BoardCandidateType(req.BoardType), req.PaymentAddress)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidArgument, service.ErrBoardCandidateExists:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		result := serializers.NewVotingBoardCandidateResp(votingBoardCandidate)
		c.JSON(http.StatusOK, serializers.Resp{Result: result})
	default:
		server.logger.Error("s.votingSvc.RegisterBoardCandidate", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (server *Server) GetCandidatesList(c *gin.Context) {
	_, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	boardQuery := c.DefaultQuery("board", "1")
	board, _ := strconv.Atoi(boardQuery)
	list, err := server.votingSvc.GetCandidatesList(board, c.Query("payment_address"))
	if err != nil {
		server.logger.Error("s.votingSvc.GetCandidatesList", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: list})
}

func (server *Server) VoteCandidateBoard(c *gin.Context) {
	var req serializers.VotingBoardCandidateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
		return
	}
	user, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	vote, err := server.votingSvc.VoteCandidateBoard(user, &req)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidBoardType, service.ErrInvalidArgument:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: vote})
	default:
		server.logger.Error("s.voting.VoteCandidateBoard", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (server *Server) CreateProposal(c *gin.Context) {
	user, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	var req serializers.VotingProposalRequest
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
		result := serializers.NewProposalDCBResp(proposal.(*models.VotingProposalDCB))
		c.JSON(http.StatusOK, serializers.Resp{Result: result})
	case 2:
		result := serializers.NewProposalGOVResp(proposal.(*models.VotingProposalGOV))
		c.JSON(http.StatusOK, serializers.Resp{Result: result})
	default:
		server.logger.Error("s.votingSvc.CreateProposal", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
}

func (server *Server) GetProposalsList(c *gin.Context) {
	vs, err := server.votingSvc.GetProposalsList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	result := make([]*serializers.ProposalResp, 0, len(vs))
	for _, v := range vs {
		switch v.(type) {
		case *models.VotingProposalDCB:
			result = append(result, serializers.NewProposalDCBResp(v.(*models.VotingProposalDCB)))
		case *models.VotingProposalGOV:
			result = append(result, serializers.NewProposalGOVResp(v.(*models.VotingProposalGOV)))
		}
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: result})
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

func (server *Server) GetBondTypes(c *gin.Context) {
	result, err := server.votingSvc.GetBondTypes()
	c.JSON(http.StatusOK, serializers.Resp{Error: err, Result: result})
}
