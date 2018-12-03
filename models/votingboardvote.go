package models

import "github.com/jinzhu/gorm"

type VotingBoardVote struct {
	gorm.Model

	VotingBoardCandidate   *VotingBoardCandidate
	VotingBoardCandidateID int

	Voter   *User `gorm:"foreignkey:VoterID"`
	VoterID uint

	BoardType int
	TxID      string
}

func (*VotingBoardVote) TableName() string {
	return "voting_board_vote"
}
