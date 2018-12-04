package models

import "github.com/jinzhu/gorm"

type ReserveContributionRequestPaymentPartyStatus int

const (
	ReserveContributionRequestPaymentPartyStatusPending ReserveContributionRequestPaymentPartyStatus = iota
	ReserveContributionRequestPaymentPartyStatusAuthorized
	ReserveContributionRequestPaymentPartyStatusSettled
	ReserveContributionRequestPaymentPartyStatusCancelled
	ReserveContributionRequestPaymentPartyStatusInvalid
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

	Status ReserveContributionRequestPaymentPartyStatus
}

func (*ReserveContributionRequestPaymentParty) TableName() string {
	return "reserve_contribution_request_payment_party"
}

func GetContributionPaymentPartyState(s string) ReserveContributionRequestPaymentPartyStatus {
	switch s {
	case "pending":
		return ReserveContributionRequestPaymentPartyStatusPending
	case "authorized":
		return ReserveContributionRequestPaymentPartyStatusAuthorized
	case "settled":
		return ReserveContributionRequestPaymentPartyStatusSettled
	case "cancelled":
		return ReserveContributionRequestPaymentPartyStatusCancelled
	default:
		return ReserveContributionRequestPaymentPartyStatusInvalid
	}
}
