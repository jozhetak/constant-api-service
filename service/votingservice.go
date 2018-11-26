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
	var candidate *models.VotingBoardCandidate
	candidate, _ = self.votingDao.FindVotingBoardCandidateByUser(*u)
	exist := false
	if candidate == nil {
		candidate = &models.VotingBoardCandidate{
			User:           u,
			PaymentAddress: paymentAddress,
		}
	} else {
		exist = true
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

	if !exist {
		candidateCreated, err := self.votingDao.CreateVotingBoardCandidate(candidate)
		if err != nil {
			return nil, err
		}
		return candidateCreated, nil
	} else {
		candidateUpdated, err := self.votingDao.UpdateVotingBoardCandidate(candidate)
		if err != nil {
			return nil, err
		}
		return candidateUpdated, nil
	}

}

func (self *VotingService) GetCandidatesList(boardType int, paymentAddress string) ([]*models.VotingBoardCandidate, error) {
	list, err := self.votingDao.Filter(&voting.VotingCandidateFilter{
		BoardType:      boardType,
		PaymentAddress: paymentAddress,
	})

	// TODO get voting number

	return list, err
}
