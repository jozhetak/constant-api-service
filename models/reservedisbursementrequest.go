package models

import "github.com/jinzhu/gorm"

type ReserveDisbursementRequestStatus int

const (
	ReserveDisbursementRequestStatusPending ReserveDisbursementRequestStatus = iota
	ReserveDisbursementRequestStatusFilled
)

type ReserveDisbursementRequest struct {
	gorm.Model

	User   *User
	UserID uint

	PartyID uint
	Status  ReserveDisbursementRequestStatus
	TxID    string
}

func (*ReserveDisbursementRequest) TableName() string {
	return "reserve_disbursement_request"
}
