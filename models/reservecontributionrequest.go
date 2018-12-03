package models

import "github.com/jinzhu/gorm"

type ReserveContributionRequestStatus int

const (
	ReserveContributionRequestStatusPending ReserveContributionRequestStatus = iota
	ReserveContributionRequestStatusFilled
)

type ReserveContributionRequest struct {
	gorm.Model

	User   *User
	UserID int

	PartyID        int
	Status         ReserveContributionRequestStatus
	TxID           string
	PaymentAddress string
}

func (*ReserveContributionRequest) TableName() string {
	return "reserve_contribution_request"
}
