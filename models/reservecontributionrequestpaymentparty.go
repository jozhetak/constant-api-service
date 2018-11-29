package models

import "github.com/jinzhu/gorm"

type ReserveContributionRequestPaymentPartyStatus int

const (
	ReserveContributionRequestPaymentPartyStatusPending ReserveContributionRequestPaymentPartyStatus = iota
)

type ReserveContributionRequestPaymentParty struct {
	gorm.Model

	ReserveContributionRequest   *ReserveContributionRequest
	ReserveContributionRequestID int

	RequestData     string
	Amount          float64
	ExtRequestData  string
	ExtResponseData string
	ExtID           string
	ExtResourceID   string
	ExtResourceType string

	Status int
}

func (*ReserveContributionRequestPaymentParty) TableName() string {
	return "reserve_contribution_request_payment_party"
}
