package models

import "github.com/jinzhu/gorm"

type ContributionRequestStatus int

const (
	ContributionRequestStatusPending ContributionRequestStatus = iota
	ContributionRequestStatusFilled
)

type ReserveContributionRequest struct {
	gorm.Model

	User   *User
	UserID int

	PartyID        int
	Status         ContributionRequestStatus
	TxID           string
	PaymentAddress string
	// TODO
}

func (*ReserveContributionRequest) TableName() string {
	return "reserve_contribution_request"
}
