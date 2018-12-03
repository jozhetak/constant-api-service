package serializers

import "github.com/ninjadotorg/constant-api-service/models"

type VotingBoardCandidateResp struct {
	User    *UserResp `json:"User"`
	DCB     string    `json:"DCB"`
	CMB     string    `json:"CMB"`
	GOV     string    `json:"GOV"`
	VoteNum int       `json:"VoteNum"`
}

func NewVotingBoardCandidateResp(data *models.VotingBoardCandidate) *VotingBoardCandidateResp {
	result := VotingBoardCandidateResp{
		GOV:     data.GOV,
		CMB:     data.CMB,
		DCB:     data.DCB,
		VoteNum: data.GetVoteNum(),
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
