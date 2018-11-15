package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

var allTables = []interface{}{(*models.User)(nil), (*models.UserVerification)(nil), (*models.UserLenderVerification)(nil), (*models.Borrow)(nil), (*models.Currency)(nil), (*models.Market)(nil), (*models.Order)(nil)}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(allTables...).Error; err != nil {
		return errors.Wrap(err, "db.AutoMigrate")
	}
	return nil
}
