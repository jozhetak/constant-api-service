package models

import "github.com/jinzhu/gorm"

type ReserveContributionRequestPaymentParty struct {
	gorm.Model

	ReserveContributionRequest   *ReserveContributionRequest
	ReserveContributionRequestID int

	// TODO
}

func (*ReserveContributionRequestPaymentParty) TableName() string {
	return "reserve_contribution_request_payment_party"
}
