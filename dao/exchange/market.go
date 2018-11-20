package exchange

import (
	"time"

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

func (e *Exchange) SymbolRates(from, to time.Time) ([]SymbolRate, error) {
	// SELECT symbol_code, sum(quantity) as volume, max(price) as high, min(price) as low FROM `exchange_orders` JOIN exchange_markets em ON em.id = exchange_orders.market_id WHERE (time >= '2018-11-19 10:52:48' AND time <= '2018-11-20 10:52:48') GROUP BY symbol_code
	var results []SymbolRate
	if err := e.db.Table("exchange_orders").Where("time >= ? AND time <= ?", from, to).Select("symbol_code, sum(quantity) as volume, max(price) as high, min(price) as low").Joins("JOIN exchange_markets em ON em.id = exchange_orders.market_id").Group("symbol_code").Scan(&results).Error; err != nil {
		return nil, errors.Wrap(err, "e.db.Table")
	}

	for i, r := range results {
		results[i].PrevPrice, _ = e.GetPriceAt(r.SymbolCode, from)
		results[i].Last, _ = e.GetLastPrice(r.SymbolCode)
	}
	return results, nil
}

func (e *Exchange) GetPriceAt(symbol string, from time.Time) (float64, error) {
	var result struct {
		Price float64
	}
	if err := e.db.Table("exchange_orders").Joins("JOIN exchange_markets em ON em.id = exchange_orders.market_id").Where("exchange_orders.updated_at >= ?", from).Where("em.symbol_code = ?", symbol).Where("status = ?", int(models.Filled)).Order("exchange_orders.updated_at ASC").Limit(1).Select("price").Scan(&result).Error; err != nil {
		return 0, errors.Wrap(err, "e.db.Table")
	}
	return result.Price, nil
}

func (e *Exchange) GetLastPrice(symbol string) (float64, error) {
	var result struct {
		Price float64
	}
	if err := e.db.Table("exchange_orders").Joins("JOIN exchange_markets em ON em.id = exchange_orders.market_id").Where("em.symbol_code = ?", symbol).Where("status = ?", int(models.Filled)).Order("exchange_orders.updated_at DESC").Limit(1).Select("price").Scan(&result).Error; err != nil {
		return 0, errors.Wrap(err, "e.db.Table")
	}
	return result.Price, nil
}

type SymbolRate struct {
	SymbolCode string
	Volume     uint
	Last       float64
	High, Low  float64
	PrevPrice  float64
}
