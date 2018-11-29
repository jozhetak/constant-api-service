package models

import "github.com/jinzhu/gorm"

type ReserveDisbursementRequestPaymentPartyStatus int

const (
	ReserveDisbursementRequestPaymentPartyStatusPending ReserveDisbursementRequestPaymentPartyStatus = iota
)

type ReserveDisbursementRequestPaymentParty struct {
	gorm.Model

	ReserveDisbursementRequest   *ReserveDisbursementRequest
	ReserveDisbursementRequestID int

	RequestData     string
	Amount          float64
	ExtRequestData  string
	ExtResponseData string
	ExtID           string
	ExtResourceID   string
	ExtResourceType string
}

func (*ReserveDisbursementRequestPaymentParty) TableName() string {
	return "reserve_disbursement_request_payment_party"
}
