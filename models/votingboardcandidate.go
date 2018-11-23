package models

import (
	_ "github.com/jinzhu/gorm"
	_ "time"
)

type VotingBoardCandidate struct {
	User   *User
	UserID int
	DCB    bool
	CMD    bool
	GOV    bool

	PaymentAddress string
}

func (*VotingBoardCandidate) TableName() string {
	return "voting_board_candidate"
}
