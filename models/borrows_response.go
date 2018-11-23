package models

import "github.com/jinzhu/gorm"

type BorrowResponse struct {
	gorm.Model
	Borrow                   Borrow
	BorrowID                 int
	ConstantLoanResponseTxID string
}

func (*BorrowResponse) TableName() string {
	return "portal_borrows_response"
}
