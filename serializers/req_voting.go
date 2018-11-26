package serializers

type VotingBoardCandidateRequest struct {
	BoardType      int    `json:"BoardType"`
	PaymentAddress string `json:"PaymentAddress"`
}
