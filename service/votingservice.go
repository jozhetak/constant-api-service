package service

import (
	"github.com/ninjadotorg/constant-api-service/dao/voting"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/pkg/errors"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"encoding/json"
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

func (self *VotingService) VoteCandidateBoard() (error) {
	// TODO call blockchain network
	// Update DB

	return nil
}

func (self *VotingService) CreateProposal(user *models.User, request *serializers.VotingProposalRequest) (models.ProposalInterface, error) {
	// TODO
	switch request.Type {
	case 1: // DCB
		{
			dcbParams := request.DCB
			// TODO call blockchain network rpc function
			// TODO waiting and check tx with blockchain network
			// TODO, call mysql
			dcbParamsStrByte, _ := json.MarshalIndent(dcbParams, "", "\t")
			proposal := &models.VotingProposalDCB{
				User: user,
				Data: string(dcbParamsStrByte),
			}
			proposalCreated, err := self.votingDao.CreateVotingProposalDCB(proposal)
			if err != nil {
				return proposalCreated, err
			}
			return proposalCreated, nil
		}
	case 2: // GOV
		{
			govParams := request.GOV
			// TODO call blockchain network rpc function
			// TODO waiting and check tx with blockchain network
			// TODO, call mysql
			govParamsStrByte, _ := json.MarshalIndent(govParams, "", "\t")
			proposal := &models.VotingProposalGOV{
				User: user,
				Data: string(govParamsStrByte),
			}
			proposalCreated, err := self.votingDao.CreateVotingProposalGOV(proposal)
			if err != nil {
				return proposalCreated, err
			}
			return proposalCreated, nil
		}
	default:
		{
			return nil, errors.New("Proposal type is wrong")
		}
	}
	return nil, nil
}

func (self *VotingService) GetProposalsList() ([]*models.ProposalInterface) {
	return nil
}

func (self *VotingService) GetProposal() (models.ProposalInterface) {
	return nil
}

func (self *VotingService) VoteProposal() (error) {
	// TODO call blockchain network
	// TODO waiting tx in block
	// Update DB
	return nil
}
