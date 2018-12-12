package models

import "github.com/jinzhu/gorm"

type VotingProposalDCB struct {
	gorm.Model

	Name string

	User   *User
	UserID int

	VotingProposalDCBVotes []*VotingProposalDCBVote
	voteNum                int

	TxID string
	Data string
}

func (v *VotingProposalDCB) SetVoteNum(num int) {
	v.voteNum = num
}

func (v *VotingProposalDCB) GetVoteNum() int {
	return v.voteNum
}

func (v *VotingProposalDCB) GetType() int {
	return 1
}

func (*VotingProposalDCB) TableName() string {
	return "voting_proposal_dcb"
}
