package serializers

import "github.com/ninjadotorg/constant-api-service/models"

type VotingBoardCandidateResp struct {
	UserID int  `json:"UserID"`
	DCB    bool `json:"DCB"`
	CMB    bool `json:"CMB"`
	GOV    bool `json:"GOV"`

	PaymentAddress string `json:"PaymentAddress"`
}

func NewVotingBoardCandidateResp(data *models.VotingBoardCandidate) *VotingBoardCandidateResp {
	result := VotingBoardCandidateResp{
		PaymentAddress: data.PaymentAddress,
		GOV:            data.GOV,
		CMB:            data.CMB,
		DCB:            data.DCB,
	}
	return &result
}
