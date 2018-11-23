package service

import (
	"github.com/ninjadotorg/constant-api-service/dao/voting"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/pkg/errors"
)

type VotingService struct {
	votingDao         *voting.VotingDao
	blockchainService *blockchain.Blockchain
}

func NewVotingService(r *voting.VotingDao, bc *blockchain.Blockchain) *VotingService {
	return &VotingService{
		votingDao:         r,
		blockchainService: bc,
	}
}

func (self *VotingService) RegisterBoardCandidate(u *models.User, boardType models.BoardCandidateType, paymentAddress string) (*models.VotingBoardCandidate, error) {
	candidate := models.VotingBoardCandidate{
		User:           u,
		PaymentAddress: paymentAddress,
	}
	switch boardType {
	case models.DCB:
		candidate.DCB = true
	case models.CMB:
		candidate.CMB = true
	case models.GOV:
		candidate.GOV = true
	default:
		return nil, errors.New("Wrong type of board")
	}

	candidateCreated, err := self.votingDao.CreateVotingBoardCandidate(&candidate)
	if err != nil {
		return nil, err
	}

	return candidateCreated, nil
}
