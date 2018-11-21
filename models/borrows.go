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
	Payment
)

type Borrow struct {
	gorm.Model
	PaymentAddress      string
	Amount              int64
	Hash                string
	CollateralTxID      string
	Collateral          string
	StartDate           time.Time
	EndDate             time.Time
	Rate                float64
	ConstantLoanTxID    string
	ConstantPaymentTxID string
	State               BorrowState `gorm:"not null;default:0"`
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
