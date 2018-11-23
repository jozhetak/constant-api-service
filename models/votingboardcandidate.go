package models

import (
	_ "github.com/jinzhu/gorm"
	_ "time"
)

type BoardCandidateType int

const (
	DCB BoardCandidateType = iota
	CMB
	GOV
)

type VotingBoardCandidate struct {
	User   *User
	UserID int
	DCB    bool
	CMB    bool
	GOV    bool

	PaymentAddress string
}

func (*VotingBoardCandidate) TableName() string {
	return "voting_board_candidate"
}