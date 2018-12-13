package models

import "github.com/jinzhu/gorm"

type VotingProposalDCB struct {
	gorm.Model

	Name string

	User   *User
	UserID int

	VotingProposalDCBVotes []*VotingProposalDCBVote

	TxID string
	Data string

	voteNum int
	isVoted bool
}

func (v *VotingProposalDCB) SetVoteNum(num int) {
	v.voteNum = num
}

func (v *VotingProposalDCB) GetVoteNum() int {
	return v.voteNum
}

func (v *VotingProposalDCB) SetIsVoted(val bool) {
	v.isVoted = val
}

func (v *VotingProposalDCB) IsVoted() bool {
	return v.isVoted
}

func (v *VotingProposalDCB) GetType() int {
	return 1
}

func (*VotingProposalDCB) TableName() string {
	return "voting_proposal_dcb"
}
