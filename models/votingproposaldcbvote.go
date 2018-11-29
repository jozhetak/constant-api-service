package models

import "github.com/jinzhu/gorm"

type VotingProposalDCBVote struct {
	gorm.Model

	VotingProposalDCB   *VotingProposalDCB
	VotingProposalDCBID int

	Voter   *User `gorm:"foreignkey:VoterID"`
	VoterID uint

	TxID string
}

func (*VotingProposalDCBVote) TableName() string {
	return "voting_proposal_dcb_vote"
}
