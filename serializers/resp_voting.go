package serializers

import (
	"time"

	"github.com/ninjadotorg/constant-api-service/models"
)

type VotingBoardCandidateResp struct {
	User         *UserResp       `json:"User"`
	DCB          string          `json:"DCB"`
	DCBAppliedAt string          `json:"DCBAppliedAt"`
	DCBBalances  *WalletBalances `json:"DCBBalances"`
	CMB          string          `json:"CMB"`
	CMBAppliedAt string          `json:"CMBAppliedAt"`
	CMBBalances  *WalletBalances `json:"CMBBalances"`
	GOV          string          `json:"GOV"`
	GOVAppliedAt string          `json:"GOVAppliedAt"`
	GOVBalances  *WalletBalances `json:"GOVBalances"`
	VoteNum      int             `json:"VoteNum"`
}

func NewVotingBoardCandidateResp(data *models.VotingBoardCandidate) *VotingBoardCandidateResp {
	if data == nil {
		return nil
	}
	result := VotingBoardCandidateResp{
		GOV:     data.GOV,
		CMB:     data.CMB,
		DCB:     data.DCB,
		VoteNum: data.GetVoteNum(),
	}
	if data.DCBAppliedAt != nil {
		result.DCBAppliedAt = data.DCBAppliedAt.Format(time.RFC3339)
	}
	if data.CMBAppliedAt != nil {
		result.CMBAppliedAt = data.CMBAppliedAt.Format(time.RFC3339)
	}
	if data.GOVAppliedAt != nil {
		result.GOVAppliedAt = data.GOVAppliedAt.Format(time.RFC3339)
	}

	result.User = NewUserResp(*data.User)
	return &result
}

type VotingBoardCandidateRespList struct {
	ListBoardCandidates []VotingBoardCandidateResp `json:"ListBoardCandidates"`
}

func NewVotingBoardCandidateListResp(data []*models.VotingBoardCandidate) *VotingBoardCandidateRespList {
	if len(data) == 0 {
		return nil
	}
	result := &VotingBoardCandidateRespList{
		ListBoardCandidates: []VotingBoardCandidateResp{},
	}
	for _, item := range data {
		temp := NewVotingBoardCandidateResp(item)
		result.ListBoardCandidates = append(result.ListBoardCandidates, *temp)
	}
	return result
}

type VotingBoardVoteResp struct {
	Voter                    *UserResp
	TxID                     string
	VotingBoardCandidateResp *VotingBoardCandidateResp
}

func NewVotingBoardVote(v *models.VotingBoardVote) *VotingBoardVoteResp {
	return &VotingBoardVoteResp{
		Voter: NewUserResp(*(v.Voter)),
		TxID:  v.TxID,
		VotingBoardCandidateResp: NewVotingBoardCandidateResp(v.VotingBoardCandidate),
	}
}

// Proposal
type ProposalResp struct {
	ID      uint      `json:"ID"`
	Name    string    `json:"Name"`
	User    *UserResp `json:"User"`
	VoteNum int       `json:"VoteNum"`

	TxID      string `json:"TxID"`
	Data      string `json:"Data"`
	CreatedAt string `json:"CreatedAt"`
}

func NewProposalDCBResp(data *models.VotingProposalDCB) *ProposalResp {
	result := &ProposalResp{
		ID:        data.ID,
		Name:      data.Name,
		Data:      data.Data,
		VoteNum:   data.GetVoteNum(),
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
	}
	result.User = NewUserResp(*data.User)
	return result
}

func NewProposalGOVResp(data *models.VotingProposalGOV) *ProposalResp {
	result := &ProposalResp{
		ID:        data.ID,
		Name:      data.Name,
		Data:      data.Data,
		VoteNum:   data.GetVoteNum(),
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
	}
	result.User = NewUserResp(*data.User)
	return result
}

// Proposal vote
type VotingProposalResp struct {
	Voter        *UserResp     `json:"Voter"`
	TxID         string        `json:"TxID"`
	ProposalResp *ProposalResp `json:"ProposalResp"`
	CreatedAt    string        `json:"CreatedAt"`
}

func NewVotingDCBProposal(data *models.VotingProposalDCBVote) *VotingProposalResp {
	return &VotingProposalResp{
		Voter:        NewUserResp(*data.Voter),
		TxID:         data.TxID,
		CreatedAt:    data.CreatedAt.Format(time.RFC3339),
		ProposalResp: NewProposalDCBResp(data.VotingProposalDCB),
	}
}

func NewVotingGOVProposal(data *models.VotingProposalGOVVote) *VotingProposalResp {
	return &VotingProposalResp{
		Voter:        NewUserResp(*data.Voter),
		TxID:         data.TxID,
		CreatedAt:    data.CreatedAt.Format(time.RFC3339),
		ProposalResp: NewProposalGOVResp(data.VotingProposalGOV),
	}
}
