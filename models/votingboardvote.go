package models

import (
	_ "github.com/jinzhu/gorm"
	_ "time"
)

type VotingBoardVote struct {
	User      *User
	UserID    int
	Voter     *User
	VoterID   uint
	BoardType int
}

func (*VotingBoardVote) TableName() string {
	return "voting_board_vote"
}
