package models

type VotingProposalDCBVote struct {
	VotingProposalDCB   *VotingProposalDCB
	VotingProposalDCBID int
	Voter               *User
	VoterID             uint
	BoardType           int
}

func (*VotingProposalDCBVote) TableName() string {
	return "voting_proposal_dcb_vote"
}
