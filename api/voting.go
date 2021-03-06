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
	candidate, err := server.votingSvc.RegisterBoardCandidate(user, models.BoardCandidateType(req.BoardType), req.PaymentAddress)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidArgument, service.ErrBoardCandidateExists:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: candidate})
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

	user, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	boardQuery := c.DefaultQuery("board_type", "1")
	board, err := strconv.Atoi(boardQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: service.ErrInvalidArgument})
	}

	list, err := server.votingSvc.GetCandidatesList(user, board, c.Query("payment_address"))
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
	case service.ErrInvalidBoardType, service.ErrInvalidArgument, service.ErrInsufficientBalance, service.ErrAlreadyVoted:
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

	var req serializers.RegisterProposalRequest
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
	var (
		boardType = c.DefaultQuery("board_type", "1")
		limit     = c.DefaultQuery("limit", "10")
		page      = c.DefaultQuery("page", "1")
	)

	user, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	vs, err := server.votingSvc.GetProposalsList(user, boardType, limit, page)
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
	var (
		id        = c.Param("id")
		boardType = c.DefaultQuery("board_type", "1")
	)

	user, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	v, err := server.votingSvc.GetProposal(id, boardType, user)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidBoardType, service.ErrInvalidProposal, service.ErrProposalNotFound:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		switch v.(type) {
		case *models.VotingProposalDCB:
			c.JSON(http.StatusOK, serializers.Resp{Result: serializers.NewProposalDCBResp(v.(*models.VotingProposalDCB))})
		case *models.VotingProposalGOV:
			c.JSON(http.StatusOK, serializers.Resp{Result: serializers.NewProposalGOVResp(v.(*models.VotingProposalGOV))})
		}
	default:
		server.logger.Error("s.voting.VoteCandidateBoard", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (server *Server) VoteProposal(c *gin.Context) {
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

	v, err := server.votingSvc.VoteProposal(user, &req)
	switch cErr := errors.Cause(err); cErr {
	case service.ErrInvalidProposal, service.ErrAlreadyVoted:
		c.JSON(http.StatusBadRequest, serializers.Resp{Error: cErr.(*service.Error)})
	case nil:
		c.JSON(http.StatusOK, serializers.Resp{Result: v})
	default:
		server.logger.Error("s.voting.VoteProposal", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
	}
}

func (server *Server) GetBondTypes(c *gin.Context) {
	result, err := server.votingSvc.GetBondTypes()
	c.JSON(http.StatusOK, serializers.Resp{Error: err, Result: result})
}

func (server *Server) GetUserCandidate(c *gin.Context) {
	user, err := server.userFromContext(c)
	if err != nil {
		server.logger.Error("s.userFromContext", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}

	candidate, err := server.votingSvc.GetUserCandidate(user)
	if err != nil {
		server.logger.Error("s.voting.VoteProposal", zap.Error(err))
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: candidate})
}

func (server *Server) GetDCBParams(c *gin.Context) {
	p, err := server.votingSvc.GetDCBParams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: p})
}

func (server *Server) GetGOVParams(c *gin.Context) {
	p, err := server.votingSvc.GetGOVParams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializers.Resp{Error: service.ErrInternalServerError})
		return
	}
	c.JSON(http.StatusOK, serializers.Resp{Result: p})
}
