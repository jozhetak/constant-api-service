package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ninjadotorg/constant-api-service/dao/voting"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/pkg/errors"
)

type VotingService struct {
	votingDao         *voting.VotingDao
	walletSvc         *WalletService
	blockchainService *blockchain.Blockchain
}

func NewVotingService(r *voting.VotingDao, bc *blockchain.Blockchain, walletSvc *WalletService) *VotingService {
	return &VotingService{
		votingDao:         r,
		walletSvc:         walletSvc,
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

func (self *VotingService) GetCandidatesList(boardType int, paymentAddress string) ([]*serializers.VotingBoardCandidateResp, error) {
	list, err := self.votingDao.Filter(&voting.VotingCandidateFilter{
		BoardType:      boardType,
		PaymentAddress: paymentAddress,
	})

	resp := make([]*serializers.VotingBoardCandidateResp, 0, len(list))
	for _, l := range list {
		r := serializers.NewVotingBoardCandidateResp(l)
		totalVote, err := self.votingDao.GetTotalVote(l.ID)
		if err != nil {
			return nil, errors.Wrap(err, "self.votingDao.GetTotalVote")
		}

		// get voting number
		r.VoteNum = totalVote

		// get balance of all token for every candidate
		wallets, err := self.walletSvc.GetCoinAndCustomTokenBalanceForPaymentAddress(l.User.PaymentAddress)
		if err != nil {
			return nil, errors.Wrap(err, "self.walletSvc.GetCoinAndCustomTokenBalanceForPaymentAddress")
		}
		r.User.WalletBalances = wallets

		resp = append(resp, r)
	}
	return resp, err
}

func (self *VotingService) VoteCandidateBoard(voter *models.User, req *serializers.VotingBoardCandidateRequest) (*serializers.VotingBoardVoteResp, error) {
	candidate, err := self.votingDao.FindCandidateByID(req.CandidateID)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.FindCandidateByID")
	}
	if candidate == nil {
		return nil, ErrInvalidArgument
	}

	var txID string
	switch models.BoardCandidateType(req.BoardType) {
	case models.DCB:
		txID, err = self.blockchainService.CreateAndSendVoteDCBBoardTransaction(voter.PrivKey, req.VoteAmount)
	case models.GOV:
		txID, err = self.blockchainService.CreateAndSendVoteGOVBoardTransaction(voter.PrivKey, req.VoteAmount)
	default:
		err = ErrInvalidBoardType
	}
	if err != nil {
		return nil, errors.Wrap(err, "self.blockchainService.CreateAndSendVoteDCBBoardTransaction")
	}

	// tx, err := GetBlockchainTxByHash(txID, 3, self.blockchainService)
	tx := &blockchain.TransactionDetail{}
	if err != nil {
		return nil, errors.Wrap(err, "GetBlockchainTxByHash")
	}
	if tx == nil {
		return nil, errors.Errorf("couldn't get tx by tx ID: %q", txID)
	}
	vote, err := self.votingDao.CreateVotingBoardVote(&models.VotingBoardVote{
		BoardType:            req.BoardType,
		Voter:                voter,
		VotingBoardCandidate: candidate,
		TxID:                 txID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.CreateVotingBoardVote")
	}

	return serializers.NewVotingBoardVote(vote), nil
}

func (self *VotingService) CreateProposal(user *models.User, request *serializers.VotingProposalRequest) (models.ProposalInterface, error) {
	// TODO
	switch request.Type {
	case 1: // DCB
		// TODO call blockchain network rpc function
		// TODO waiting and check tx with blockchain network
		params, _ := json.Marshal(request.DCB)
		proposal := &models.VotingProposalDCB{
			User: user,
			Data: string(params),
			TxID: "", // get tx above
		}
		proposalCreated, err := self.votingDao.CreateVotingProposalDCB(proposal)
		if err != nil {
			return proposalCreated, err
		}
		return proposalCreated, nil
	case 2: // GOV
		// TODO call blockchain network rpc function
		// TODO waiting and check tx with blockchain network
		params, _ := json.Marshal(request.GOV)
		proposal := &models.VotingProposalGOV{
			User: user,
			Data: string(params),
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

func (self *VotingService) getDCBProposals(limit, page *int) ([]models.ProposalInterface, error) {
	vs, err := self.votingDao.GetDCBProposals(limit, page)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.GetDCBProposals")
	}
	ret := make([]models.ProposalInterface, 0, len(vs))
	for _, v := range vs {
		totalVote, err := self.votingDao.GetProposalDCBVote(v.ID)
		if err != nil {
			return nil, errors.Wrap(err, "self.votingDao.GetProposalDCBVote")
		}
		v.SetVoteNum(totalVote)

		ret = append(ret, v)
	}
	return ret, nil
}

func (self *VotingService) getGOVProposals(limit, page *int) ([]models.ProposalInterface, error) {
	vs, err := self.votingDao.GetGOVProposals(limit, page)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.GetDCBProposals")
	}
	ret := make([]models.ProposalInterface, 0, len(vs))
	for _, v := range vs {
		totalVote, err := self.votingDao.GetProposalGOVVote(v.ID)
		if err != nil {
			return nil, errors.Wrap(err, "self.votingDao.GetProposalGOVVote")
		}
		v.SetVoteNum(totalVote)

		ret = append(ret, v)
	}
	return ret, nil
}

func (self *VotingService) GetProposalsList(boardType, limit, page string) ([]models.ProposalInterface, error) {
	l, p, err := parsePaginationQuery(limit, page)
	if err != nil {
		return nil, errors.Wrap(err, "parsePaginationQuery")
	}
	typ, err := strconv.Atoi(boardType)
	if err != nil {
		return nil, ErrInvalidBoardType
	}
	switch models.BoardCandidateType(typ) {
	case models.DCB:
		return self.getDCBProposals(&l, &p)
	case models.GOV:
		return self.getGOVProposals(&l, &p)
	default:
		return nil, ErrInvalidBoardType
	}
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
