package models

import "github.com/jinzhu/gorm"

type DisbursementRequestStatus int

const (
	DisbursementRequestStatusPending        DisbursementRequestStatus = iota
	DisbursementRequestStatusAuthorinzation
	DisbursementRequestStatusSettlement
)

type ReserveDisbursementRequest struct {
	gorm.Model

	User   *User
	UserID int

	PartyID int
	Status  DisbursementRequestStatus
	TxID    string
}

func (*ReserveDisbursementRequest) TableName() string {
	return "reserve_disbursement_request"
}
