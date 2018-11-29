package serializers

import "github.com/ninjadotorg/constant-api-service/models"

type VotingBoardCandidateResp struct {
	User    *UserResp `json:"User"`
	UserID  int       `json:"UserID"`
	DCB     string    `json:"DCB"`
	CMB     string    `json:"CMB"`
	GOV     string    `json:"GOV"`
	VoteNum int       `json:"VoteNum"`

	PaymentAddress string `json:"PaymentAddress"`
}

func NewVotingBoardCandidateResp(data *models.VotingBoardCandidate) *VotingBoardCandidateResp {
	result := VotingBoardCandidateResp{
		GOV:     data.GOV,
		CMB:     data.CMB,
		DCB:     data.DCB,
		VoteNum: data.GetVoteNum(),
	}

	result.UserID = data.UserID
	result.User = NewUserResp(*data.User)
	return &result
}

type VotingBoardCandidateRespList struct {
	ListBoardCandidates []VotingBoardCandidateResp `json:"ListBoardCandidates"`
}

func NewVotingBoardCandidateListResp(data []*models.VotingBoardCandidate) *VotingBoardCandidateRespList {
	result := &VotingBoardCandidateRespList{
		ListBoardCandidates: []VotingBoardCandidateResp{},
	}
	for _, item := range data {
		temp := NewVotingBoardCandidateResp(item)
		result.ListBoardCandidates = append(result.ListBoardCandidates, *temp)
	}
	return result
}

// Proposal
type ProposalResp struct {
	User    UserResp `json:"User"`
	UserID  int
	VoteNum int `json:"VoteNum"`

	TxID string `json:"TxID"`
	Data string `json:"Data"`
}

func NewProposalDCBResp(data *models.VotingProposalDCB) *ProposalResp {
	result := &ProposalResp{
		Data:    data.Data,
		VoteNum: data.GetVoteNum(),
	}
	result.UserID = data.UserID
	result.User = *(NewUserResp(*data.User))
	return result
}

func NewProposalGOVResp(data *models.VotingProposalGOV) *ProposalResp {
	result := &ProposalResp{
		Data:    data.Data,
		VoteNum: data.GetVoteNum(),
	}
	result.UserID = data.UserID
	result.User = *(NewUserResp(*data.User))
	return result
}
