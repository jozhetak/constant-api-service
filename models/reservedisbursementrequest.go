package models

import "github.com/jinzhu/gorm"

type ReserveDisbursementRequest struct {
	gorm.Model

	User   *User
	UserID int

	// TODO
}

func (*ReserveDisbursementRequest) TableName() string {
	return "reserve_disbursement_request"
}
