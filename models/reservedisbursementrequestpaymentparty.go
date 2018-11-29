package models

import "github.com/jinzhu/gorm"

type ReserveDisbursementRequestPaymentParty struct {
	gorm.Model

	ReserveDisbursementRequest   *ReserveDisbursementRequest
	ReserveDisbursementRequestID int

	// TODO
}

func (*ReserveDisbursementRequestPaymentParty) TableName() string {
	return "reserve_disbursement_request_payment_party"
}
