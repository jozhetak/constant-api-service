package models

type BoardCandidateType int

const (
	Invalid BoardCandidateType = iota
	DCB
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

	VotingBoardVotes []*VotingBoardVote

	voteNum int
}

func (self VotingBoardCandidate) SetVoteNum(num int) {
	self.voteNum = num
}

func (self VotingBoardCandidate) GetVoteNum() int {
	return self.voteNum
}

func (*VotingBoardCandidate) TableName() string {
	return "voting_board_candidate"
}
