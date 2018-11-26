package models

import (
	_ "github.com/jinzhu/gorm"
	_ "time"
)

type VotingProposalGOV struct {
	User   *User
	UserID int
	Data   string
}

func (*VotingProposalGOV) TableName() string {
	return "voting_proposal_gov"
}
