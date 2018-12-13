package service

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ninjadotorg/constant-api-service/dao/voting"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/pkg/errors"
)

const (
	DCBCandidateToken = "0000000000000000000000000000000000000000000000000000000000000001"
	GOVCandidateToken = "0000000000000000000000000000000000000000000000000000000000000002"
	CMBCandidateToken = "0000000000000000000000000000000000000000000000000000000000000003"

	DCBProposalToken = "0000000000000000000000000000000000000000000000000000000000000005"
	GOVProposalToken = "0000000000000000000000000000000000000000000000000000000000000006"
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

func (self *VotingService) RegisterBoardCandidate(u *models.User, boardType models.BoardCandidateType, paymentAddress string) (*serializers.VotingBoardCandidateResp, error) {
	candidate, err := self.votingDao.FindVotingBoardCandidateByUser(*u)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.FindVotingBoardCandidateByUser")
	}

	var exist bool
	if candidate == nil {
		candidate = &models.VotingBoardCandidate{User: u}
	} else {
		exist = true
	}
	if paymentAddress == "" {
		paymentAddress = u.PaymentAddress
	}

	now := time.Now().UTC()
	switch boardType {
	case models.DCB:
		if candidate.DCB != "" {
			return nil, ErrBoardCandidateExists
		}
		candidate.DCB = paymentAddress
		candidate.DCBAppliedAt = &now
	case models.CMB:
		if candidate.CMB != "" {
			return nil, ErrBoardCandidateExists
		}
		candidate.CMB = paymentAddress
		candidate.CMBAppliedAt = &now
	case models.GOV:
		if candidate.GOV != "" {
			return nil, ErrBoardCandidateExists
		}
		candidate.GOV = paymentAddress
		candidate.GOVAppliedAt = &now
	default:
		return nil, ErrInvalidArgument
	}

	if exist {
		candidate, err = self.votingDao.UpdateVotingBoardCandidate(candidate)
		err = errors.Wrap(err, "self.votingDao.UpdateVotingBoardCandidate")
	} else {
		candidate, err = self.votingDao.CreateVotingBoardCandidate(candidate)
		err = errors.Wrap(err, "self.votingDao.CreateVotingBoardCandidate")
	}
	if err != nil {
		return nil, err
	}

	resp := serializers.NewVotingBoardCandidateResp(candidate)
	if err := self.GetCandidateBalances(resp); err != nil {
		return nil, errors.Wrap(err, "self.GetCandidateBalances")
	}

	return resp, nil
}

func (self *VotingService) GetCandidatesList(user *models.User, boardType int, paymentAddress string) ([]*serializers.VotingBoardCandidateResp, error) {
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
		for _, v := range l.VotingBoardVotes {
			bt := models.BoardCandidateType(v.BoardType)
			switch bt {
			case models.DCB:
				r.IsVotedDCB = (l.DCB != "")
			case models.GOV:
				r.IsVotedGOV = (l.GOV != "")
			case models.CMB:
				r.IsVotedCMB = (l.CMB != "")
			}
		}

		if err := self.GetCandidateBalances(r); err != nil {
			return nil, errors.Wrap(err, "self.GetCandidateBalances")
		}

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
	for _, v := range candidate.VotingBoardVotes {
		if voter.ID == v.Voter.ID {
			return nil, ErrAlreadyVoted
		}
	}

	var tokenID string
	switch models.BoardCandidateType(req.BoardType) {
	case models.DCB:
		tokenID = DCBCandidateToken
	case models.GOV:
		tokenID = GOVCandidateToken
	case models.CMB:
		tokenID = CMBCandidateToken
	default:
		return nil, ErrInvalidBoardType
	}
	if err := self.validateBalance(voter, tokenID, req.VoteAmount); err != nil {
		return nil, errors.Wrap(err, "self.validateBalance")
	}

	var txID string
	switch models.BoardCandidateType(req.BoardType) {
	case models.DCB:
		// txID, err = self.blockchainService.CreateAndSendVoteDCBBoardTransaction(voter.PrivKey, req.VoteAmount)
	case models.GOV:
		// txID, err = self.blockchainService.CreateAndSendVoteGOVBoardTransaction(voter.PrivKey, req.VoteAmount)
	default:
		err = ErrInvalidBoardType
	}
	// if err != nil {
	//         return nil, errors.Wrap(err, "self.blockchainService.CreateAndSendVoteDCBBoardTransaction")
	// }

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

func (self *VotingService) CreateProposal(user *models.User, request *serializers.RegisterProposalRequest) (models.ProposalInterface, error) {
	switch models.BoardCandidateType(request.Type) {
	case models.DCB:
		// TODO call blockchain network rpc function
		// TODO waiting and check tx with blockchain network
		params, _ := json.Marshal(request.DCB)
		proposal := &models.VotingProposalDCB{
			Name: request.Name,
			User: user,
			Data: string(params),
			TxID: "", // get tx above
		}
		p, err := self.votingDao.CreateVotingProposalDCB(proposal)
		if err != nil {
			return nil, errors.Wrap(err, "self.votingDao.CreateVotingProposalDCB")
		}
		return p, nil
	case models.GOV:
		// TODO call blockchain network rpc function
		// TODO waiting and check tx with blockchain network
		params, _ := json.Marshal(request.GOV)
		proposal := &models.VotingProposalGOV{
			Name: request.Name,
			User: user,
			Data: string(params),
			TxID: "", // get tx above
		}
		p, err := self.votingDao.CreateVotingProposalGOV(proposal)
		if err != nil {
			return nil, errors.Wrap(err, "self.votingDao.CreateVotingProposalGOV")
		}
		return p, nil
	default:
		return nil, errors.Errorf("unsupported proposal type: %v", request.Type)
	}
}

func (self *VotingService) isVoted(votes interface{}, user *models.User) bool {
	switch v := votes.(type) {
	case []*models.VotingProposalDCBVote:
		for _, vote := range v {
			if vote.Voter.ID == user.ID {
				return true
			}
		}
	case []*models.VotingProposalGOVVote:
		for _, vote := range v {
			if vote.Voter.ID == user.ID {
				return true
			}
		}
	}
	return false
}

func (self *VotingService) getDCBProposals(user *models.User, limit, page *int) ([]models.ProposalInterface, error) {
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
		v.SetIsVoted(self.isVoted(v.VotingProposalDCBVotes, user))

		ret = append(ret, v)
	}
	return ret, nil
}

func (self *VotingService) getGOVProposals(user *models.User, limit, page *int) ([]models.ProposalInterface, error) {
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
		v.SetIsVoted(self.isVoted(v.VotingProposalGOVVotes, user))

		ret = append(ret, v)
	}
	return ret, nil
}

func (self *VotingService) GetProposalsList(user *models.User, boardType, limit, page string) ([]models.ProposalInterface, error) {
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
		return self.getDCBProposals(user, &l, &p)
	case models.GOV:
		return self.getGOVProposals(user, &l, &p)
	default:
		return nil, ErrInvalidBoardType
	}
}

func (self *VotingService) getDCBProposal(id int, user *models.User) (models.ProposalInterface, error) {
	v, err := self.votingDao.GetDCBProposal(id)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.GetDCBProposal")
	}
	if v == nil {
		return nil, ErrProposalNotFound
	}
	total, err := self.votingDao.GetProposalDCBVote(v.ID)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.GetProposalDCBVote")
	}
	v.SetVoteNum(total)
	v.SetIsVoted(self.isVoted(v.VotingProposalDCBVotes, user))

	return v, nil
}

func (self *VotingService) getGOVProposal(id int, user *models.User) (models.ProposalInterface, error) {
	v, err := self.votingDao.GetGOVProposal(id)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.GetGOVProposal")
	}
	if v == nil {
		return nil, ErrProposalNotFound
	}
	total, err := self.votingDao.GetProposalGOVVote(v.ID)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.GetProposalGOVVote")
	}
	v.SetVoteNum(total)
	v.SetIsVoted(self.isVoted(v.VotingProposalGOVVotes, user))

	return v, nil
}

func (self *VotingService) GetProposal(id, boardType string, user *models.User) (models.ProposalInterface, error) {
	idI, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrInvalidBoardType
	}
	typ, err := strconv.Atoi(boardType)
	if err != nil {
		return nil, ErrInvalidProposal
	}
	switch models.BoardCandidateType(typ) {
	case models.DCB:
		return self.getDCBProposal(idI, user)
	case models.GOV:
		return self.getGOVProposal(idI, user)
	default:
		return nil, ErrInvalidBoardType
	}
}

func (self *VotingService) VoteProposal(u *models.User, req *serializers.VotingProposalRequest) (*serializers.VotingProposalResp, error) {
	// TODO call blockchain network
	// TODO waiting tx in block
	// Update DB
	switch models.BoardCandidateType(req.BoardType) {
	case models.DCB:
		if err := self.validateBalance(u, DCBProposalToken, req.VoteAmount); err != nil {
			return nil, errors.Wrap(err, "self.validateBalance")
		}
		p, err := self.votingDao.GetDCBProposal(req.ProposalID)
		if err != nil {
			return nil, errors.Wrap(err, "self.votingDao.GetDCBProposal")
		}
		if p == nil {
			return nil, ErrInvalidProposal
		}
		for _, v := range p.VotingProposalDCBVotes {
			if u.ID == v.Voter.ID {
				return nil, ErrAlreadyVoted
			}
		}
		v, err := self.votingDao.CreateVotingProposalDCBVote(&models.VotingProposalDCBVote{
			Voter:             u,
			VotingProposalDCB: p,
			TxID:              "", // get from above call
		})
		if err != nil {
			return nil, errors.Wrap(err, "self.votingDao.CreateVotingProposalDCBVote")
		}
		return serializers.NewVotingDCBProposal(v), nil
	case models.GOV:
		if err := self.validateBalance(u, GOVProposalToken, req.VoteAmount); err != nil {
			return nil, errors.Wrap(err, "self.validateBalance")
		}
		p, err := self.votingDao.GetGOVProposal(req.ProposalID)
		if err != nil {
			return nil, errors.Wrap(err, "self.votingDao.GetDCBProposal")
		}
		if p == nil {
			return nil, ErrInvalidProposal
		}
		for _, v := range p.VotingProposalGOVVotes {
			if u.ID == v.Voter.ID {
				return nil, ErrAlreadyVoted
			}
		}
		v, err := self.votingDao.CreateVotingProposalGOVVote(&models.VotingProposalGOVVote{
			Voter:             u,
			VotingProposalGOV: p,
			TxID:              "", // get from above call
		})
		if err != nil {
			return nil, errors.Wrap(err, "self.votingDao.CreateVotingProposalDCBVote")
		}
		return serializers.NewVotingGOVProposal(v), nil
	default:
		return nil, ErrInvalidBoardType
	}
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

func (self *VotingService) GetUserCandidate(u *models.User) (*serializers.VotingBoardCandidateResp, error) {
	c, err := self.votingDao.FindCandidateByUser(u)
	if err != nil {
		return nil, errors.Wrap(err, "self.votingDao.FindCandidateByUser")
	}
	if c == nil {
		return nil, nil
	}

	resp := serializers.NewVotingBoardCandidateResp(c)
	// uncomment this to get balances for candidate
	// if err := self.GetCandidateBalances(resp); err != nil {
	//         return nil, errors.Wrap(err, "self.GetCandidateBalances")
	// }
	return resp, nil
}

func (self *VotingService) GetCandidateBalances(resp *serializers.VotingBoardCandidateResp) error {
	if resp.CMB != "" {
		wallets, err := self.walletSvc.GetCoinAndCustomTokenBalanceForPaymentAddress(resp.CMB)
		if err != nil {
			return errors.Wrap(err, "self.blockchainService.GetCoinAndCustomTokenBalanceForPaymentAddress")
		}
		resp.CMBBalances = wallets
	}
	if resp.GOV != "" {
		wallets, err := self.walletSvc.GetCoinAndCustomTokenBalanceForPaymentAddress(resp.GOV)
		if err != nil {
			return errors.Wrap(err, "self.blockchainService.GetCoinAndCustomTokenBalanceForPaymentAddress")
		}
		resp.GOVBalances = wallets
	}
	if resp.DCB != "" {
		wallets, err := self.walletSvc.GetCoinAndCustomTokenBalanceForPaymentAddress(resp.DCB)
		if err != nil {
			return errors.Wrap(err, "self.blockchainService.GetCoinAndCustomTokenBalanceForPaymentAddress")
		}
		resp.DCBBalances = wallets
	}
	return nil
}

func (self *VotingService) validateBalance(u *models.User, tokenID string, amount uint64) error {
	balances, err := self.walletSvc.GetCoinAndCustomTokenBalanceForUser(u)
	if err != nil {
		return errors.Wrap(err, "e.w.GetCoinAndCustomTokenBalance")
	}
	for _, b := range balances.ListBalances {
		if b.TokenID == tokenID {
			if amount > b.AvailableBalance {
				return ErrInsufficientBalance
			}
		}
	}
	return ErrInsufficientBalance
}
