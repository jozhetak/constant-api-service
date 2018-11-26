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
}

func (*VotingProposalDCB) TableName() string {
	return "voting_proposal_dcb"
}
