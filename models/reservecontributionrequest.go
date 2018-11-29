package models

import "github.com/jinzhu/gorm"

type ReserveContributionRequest struct {
	gorm.Model

	User   *User
	UserID int

	// TODO
}

func (*ReserveContributionRequest) TableName() string {
	return "reserve_contribution_request"
}
