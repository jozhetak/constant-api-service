package service

import (
	"encoding/json"
	"fmt"

	"github.com/ninjadotorg/constant-api-service/dao/voting"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
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
	candidate, _ := self.votingDao.FindVotingBoardCandidateByUser(*u)
	var exist bool

	if candidate == nil {
		candidate = &models.VotingBoardCandidate{User: u}
	} else {
		exist = true
	}

	switch boardType {
	case models.DCB:
		if candidate.DCB != "" {
			return nil, ErrBoardCandidateExists
		}
		candidate.DCB = paymentAddress
	case models.CMB:
		if candidate.CMB != "" {
			return nil, ErrBoardCandidateExists
		}
		candidate.CMB = paymentAddress
	case models.GOV:
		if candidate.GOV != "" {
			return nil, ErrBoardCandidateExists
		}
		candidate.GOV = paymentAddress
	default:
		return nil, ErrInvalidArgument
	}

	if exist {
		candidateUpdated, err := self.votingDao.UpdateVotingBoardCandidate(candidate)
		if err != nil {
			return nil, errors.Wrap(err, "self.votingDao.UpdateVotingBoardCandidate")
		}
		return candidateUpdated, nil
	}
	candidateCreated, err := self.votingDao.CreateVotingBoardCandidate(candidate)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.CreateVotingBoardCandidate")
	}
	fmt.Printf("candidateCreated = %+v\n", candidateCreated)
	return candidateCreated, nil
}

func (self *VotingService) GetCandidatesList(boardType int, paymentAddress string) ([]*models.VotingBoardCandidate, error) {
	list, err := self.votingDao.Filter(&voting.VotingCandidateFilter{
		BoardType:      boardType,
		PaymentAddress: paymentAddress,
	})

	// TODO get voting number

	// TODO get balance of all token for every candidate

	return list, err
}

func (self *VotingService) VoteCandidateBoard() error {
	// TODO call blockchain network
	// Update DB

	return nil
}

func (self *VotingService) CreateProposal(user *models.User, request *serializers.VotingProposalRequest) (models.ProposalInterface, error) {
	// TODO
	switch request.Type {
	case 1: // DCB
		dcbParams := request.DCB
		// TODO call blockchain network rpc function
		// TODO waiting and check tx with blockchain network
		dcbParamsStrByte, _ := json.MarshalIndent(dcbParams, "", "\t")
		proposal := &models.VotingProposalDCB{
			User: user,
			Data: string(dcbParamsStrByte),
			TxID: "", // get tx above
		}
		proposalCreated, err := self.votingDao.CreateVotingProposalDCB(proposal)
		if err != nil {
			return proposalCreated, err
		}
		return proposalCreated, nil
	case 2: // GOV
		govParams := request.GOV
		// TODO call blockchain network rpc function
		// TODO waiting and check tx with blockchain network
		govParamsStrByte, _ := json.MarshalIndent(govParams, "", "\t")
		proposal := &models.VotingProposalGOV{
			User: user,
			Data: string(govParamsStrByte),
			TxID: "", // get tx above
		}
		proposalCreated, err := self.votingDao.CreateVotingProposalGOV(proposal)
		if err != nil {
			return proposalCreated, err
		}
		return proposalCreated, nil
	default:
		return nil, errors.Errorf("unsupported proposal type: %v", request.Type)
	}
}

func (self *VotingService) GetProposalsList() ([]models.ProposalInterface, error) {
	vs, err := self.votingDao.GetDCBProposals(nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.GetDCBProposals")
	}
	ret := make([]models.ProposalInterface, 0, len(vs))
	for _, v := range vs {
		ret = append(ret, v)
	}
	return ret, nil
}

func (self *VotingService) GetProposal() models.ProposalInterface {
	return nil
}

func (self *VotingService) VoteProposal() error {
	// TODO call blockchain network
	// TODO waiting tx in block
	// Update DB
	return nil
}

func (self *VotingService) GetBondTypes() (interface{}, error) {
	return self.blockchainService.GetBondTypes()
}

func (self *VotingService) GetGOVParams() (*serializers.VotingProposalGOVRequest, error) {
	result := serializers.VotingProposalGOVRequest{}
	blockchainData, err := self.blockchainService.GetGOVParams()
	if err != nil {
		return nil, err
	}
	temp, err := json.MarshalIndent(blockchainData, "", "\t")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(temp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *VotingService) GetDCBParams() (*serializers.VotingProposalDCBRequest, error) {
	result := serializers.VotingProposalDCBRequest{}
	blockchainData, err := self.blockchainService.GetDCBParams()
	if err != nil {
		return nil, err
	}
	temp, err := json.MarshalIndent(blockchainData, "", "\t")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(temp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
