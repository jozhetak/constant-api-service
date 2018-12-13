package models

import "github.com/jinzhu/gorm"

type VotingProposalGOV struct {
	gorm.Model

	Name string

	User   *User
	UserID int

	VotingProposalGOVVotes []*VotingProposalGOVVote

	TxID string
	Data string

	voteNum int
	isVoted bool
}

func (v *VotingProposalGOV) SetVoteNum(num int) {
	v.voteNum = num
}

func (v *VotingProposalGOV) GetVoteNum() int {
	return v.voteNum
}

func (v *VotingProposalGOV) SetIsVoted(val bool) {
	v.isVoted = val
}

func (v *VotingProposalGOV) IsVoted() bool {
	return v.isVoted
}

func (v *VotingProposalGOV) GetType() int {
	return 2
}

func (*VotingProposalGOV) TableName() string {
	return "voting_proposal_gov"
}
