package exchange

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

func (e *Exchange) ListMarkets(base string) ([]*models.Market, error) {
	var currencies []*models.Market

	query := e.db.Order("id ASC")
	if base != "" {
		query = query.Where("base_currency = ?", base)
	}
	if err := query.Find(&currencies).Error; err != nil {
		return nil, errors.Wrap(err, "c.db.Where.Find")
	}
	return currencies, nil
}

func (e *Exchange) FindMarketBySymbol(s string) (*models.Market, error) {
	var m models.Market
	if err := e.db.Where("symbol = ?", s).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}
	return &m, nil
}
