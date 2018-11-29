package models

import "github.com/jinzhu/gorm"

type BoardCandidateType int

const (
	Invalid BoardCandidateType = iota
	DCB
	CMB
	GOV
)

type VotingBoardCandidate struct {
	gorm.Model

	User   *User
	UserID int

	DCB string
	CMB string
	GOV string

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
