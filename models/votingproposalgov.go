package models

type VotingProposalGOV struct {
	User   *User
	UserID int
	Data   string

	VotingProposalGOVVotes []*VotingProposalGOVVote
	voteNum                int
}

func (self VotingProposalGOV) SetVoteNum(num int) {
	self.voteNum = num
}

func (self VotingProposalGOV) GetVoteNum() int {
	return self.voteNum
}

func (self VotingProposalGOV) GetType() int {
	return 2
}

func (*VotingProposalGOV) TableName() string {
	return "voting_proposal_gov"
}
