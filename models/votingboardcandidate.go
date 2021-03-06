package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BoardCandidateType int

const (
	Invalid BoardCandidateType = iota
	DCB
	GOV
	CMB
)

type VotingBoardCandidate struct {
	gorm.Model

	User   *User
	UserID int

	DCB          string
	DCBAppliedAt *time.Time

	CMB          string
	CMBAppliedAt *time.Time

	GOV          string
	GOVAppliedAt *time.Time

	VotingBoardVotes []*VotingBoardVote

	voteNum int
}

func (v *VotingBoardCandidate) SetVoteNum(num int) {
	v.voteNum = num
}

func (v *VotingBoardCandidate) GetVoteNum() int {
	return v.voteNum
}

func (*VotingBoardCandidate) TableName() string {
	return "voting_board_candidate"
}
