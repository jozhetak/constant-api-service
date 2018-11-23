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
	UserID                   int
	User                     *User
	PaymentAddress           string
	LoanAmount               int64
	KeyDigest                string
	LoanID                   string
	CollateralType           string
	CollateralAmount         string
	StartDate                time.Time
	EndDate                  time.Time
	InterestRate             int64
	Maturity                 int64
	LiquidationStart         int64
	ConstantLoanRequestTxID  string
	ConstantLoanResponseTxID string
	ConstantLoanWithdrawTxID string
	ConstantLoanPaymentTxID  string
	State                    BorrowState `gorm:"not null;default:0"`
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
