package models

import "github.com/jinzhu/gorm"

type ReserveDisbursementRequestPaymentPartyStatus int

const (
	ReserveDisbursementRequestPaymentPartyStatusPending ReserveDisbursementRequestPaymentPartyStatus = iota
	ReserveDisbursementRequestPaymentPartyStatusAuthorized
	ReserveDisbursementRequestPaymentPartyStatusSettled
	ReserveDisbursementRequestPaymentPartyStatusCancelled
	ReserveDisbursementRequestPaymentPartyStatusInvalid
)

type ReserveDisbursementRequestPaymentParty struct {
	gorm.Model

	ReserveDisbursementRequest   *ReserveDisbursementRequest
	ReserveDisbursementRequestID int

	RequestData     string `gorm:"type:text"`
	Amount          float64
	ExtRequestData  string `gorm:"type:text"`
	ExtResponseData string `gorm:"type:text"`
	ExtID           string
	ExtResourceID   string
	ExtResourceType string

	Status ReserveDisbursementRequestPaymentPartyStatus
}

func (*ReserveDisbursementRequestPaymentParty) TableName() string {
	return "reserve_disbursement_request_payment_party"
}

func GetDisbursementPaymentPartyState(s string) ReserveDisbursementRequestPaymentPartyStatus {
	switch s {
	case "pending":
		return ReserveDisbursementRequestPaymentPartyStatusPending
	case "authorized":
		return ReserveDisbursementRequestPaymentPartyStatusAuthorized
	case "settled":
		return ReserveDisbursementRequestPaymentPartyStatusSettled
	case "cancelled":
		return ReserveDisbursementRequestPaymentPartyStatusCancelled
	default:
		return ReserveDisbursementRequestPaymentPartyStatusInvalid
	}
}
