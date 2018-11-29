package models

import "github.com/jinzhu/gorm"

type VotingProposalGOVVote struct {
	gorm.Model

	VotingProposalGOV   *VotingProposalGOV
	VotingProposalGOVID int

	Voter   *User `gorm:"foreignkey:VoterID"`
	VoterID uint

	TxID string
}

func (*VotingProposalGOVVote) TableName() string {
	return "voting_proposal_gov_vote"
}
