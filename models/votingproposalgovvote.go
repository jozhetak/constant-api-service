package models

import (
	_ "github.com/jinzhu/gorm"
	_ "time"
)

type VotingProposalGOVVote struct {
	VotingProposalGOV   *VotingProposalGOV
	VotingProposalGOVID int
	Voter               *User
	VoterID             uint
	BoardType           int
}

func (*VotingProposalGOVVote) TableName() string {
	return "voting_proposal_gov_vote"
}
