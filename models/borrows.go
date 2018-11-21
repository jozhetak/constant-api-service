package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BorrowState int

const (
	InvalidState BorrowState = iota
	Pending
	Approved
	Rejected
)

type Borrow struct {
	gorm.Model

	User   *User
	UserID int

	Amount         int64
	Hash           string
	TxID           string
	PaymentAddress string
	StartDate      time.Time
	EndDate time.Time
	Collateral     string
	Rate           float64
	State          BorrowState `gorm:"not null;default:0"`
}

func (*Borrow) TableName() string {
	return "portal_borrows"
}

func (b BorrowState) String() string {
	return [...]string{"invalid", "pending", "approved", "rejected"}[b]
}

func GetBorrowState(s string) BorrowState {
	switch s {
	case "pending":
		return Pending
	case "approved":
		return Approved
	case "rejected":
		return Rejected
	default:
		return InvalidState
	}
}
