package models

import "github.com/jinzhu/gorm"

type VotingProposalGOV struct {
	gorm.Model

	User   *User
	UserID int

	VotingProposalGOVVotes []*VotingProposalGOVVote
	voteNum                int

	TxID string
	Data string
}

func (v *VotingProposalGOV) SetVoteNum(num int) {
	v.voteNum = num
}

func (v *VotingProposalGOV) GetVoteNum() int {
	return v.voteNum
}

func (v *VotingProposalGOV) GetType() int {
	return 2
}

func (*VotingProposalGOV) TableName() string {
	return "voting_proposal_gov"
}
