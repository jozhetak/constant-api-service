package models

import (
	_ "github.com/jinzhu/gorm"
	_ "time"
)

type VotingProposalDCB struct {
	User   *User
	UserID int
	Data   string

	VotingProposalDCBVotes []VotingProposalDCBVote
	voteNum                int
}

func (self VotingProposalDCB) SetVoteNum(num int) {
	self.voteNum = num
}

func (self VotingProposalDCB) GetVoteNum() int {
	return self.voteNum
}

func (self VotingProposalDCB) GetType() int {
	return 1
}

func (*VotingProposalDCB) TableName() string {
	return "voting_proposal_dcb"
}
