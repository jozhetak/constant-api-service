package models

import "github.com/jinzhu/gorm"

type DisbursementRequestStatus int

const (
	DisbursementRequestStatusPending DisbursementRequestStatus = iota
	DisbursementRequestStatusFilled
)

type ReserveDisbursementRequest struct {
	gorm.Model

	User   *User
	UserID uint

	PartyID uint
	Status  DisbursementRequestStatus
	TxID    string
}

func (*ReserveDisbursementRequest) TableName() string {
	return "reserve_disbursement_request"
}
