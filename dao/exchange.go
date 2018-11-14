package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/pkg/errors"
)

type Exchange struct {
	db *gorm.DB
}

func NewExchange(db *gorm.DB) *Exchange {
	return &Exchange{db}
}

func (e *Exchange) ListMarkets(base string) ([]*models.Market, error) {
	var currencies []*models.Market

	query := e.db.Order("id ASC")
	if base != "" {
		query = query.Where("base = ?", base)
	}
	if err := query.Find(&currencies).Error; err != nil {
		return nil, errors.Wrap(err, "c.db.Where.Find")
	}
	return currencies, nil
}

func (e *Exchange) CreateOrder(o *models.Order) (*models.Order, error) {
	if err := e.db.Create(o).Error; err != nil {
		return nil, errors.Wrap(err, "e.db.Create")
	}
	return o, nil
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
