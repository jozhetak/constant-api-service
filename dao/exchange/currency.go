package exchange

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

func (e *Exchange) FindCurrencyByToken(token string) (*models.Currency, error) {
	var c models.Currency
	if err := e.db.Where("token = ?", token).Find(&c).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "b.db.First")
	}
	return &c, nil
}
