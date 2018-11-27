package exchange

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

func (e *Exchange) ListMarkets(base string) ([]*models.Market, error) {
	var markets []*models.Market

	query := e.db.Order("id ASC")
	if base != "" {
		query = query.Where("base_currency = ?", base)
	}
	if err := query.Preload("BaseCurrency").Preload("QuoteCurrency").Find(&markets).Error; err != nil {
		return nil, errors.Wrap(err, "c.db.Where.Find")
	}
	return markets, nil
}

func (e *Exchange) FindMarketBySymbol(s string) (*models.Market, error) {
	var m models.Market
	if err := e.db.Where("symbol_code = ?", s).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "e.db.Where")
	}
	return &m, nil
}

func (e *Exchange) SymbolRates(from, to time.Time) ([]SymbolRate, error) {
	var results []SymbolRate

	if err := e.db.Table("exchange_orders").Joins("JOIN exchange_markets em ON em.id = exchange_orders.market_id").Where("time >= ? AND time <= ?", from, to).Select("symbol_code, sum(quantity) as volume, max(price) as high, min(price) as low").Group("symbol_code").Scan(&results).Error; err != nil {
		return nil, errors.Wrap(err, "e.db.Table")
	}

	// only get filled orders
	status := models.Filled
	for i, r := range results {
		results[i].PrevPrice, _ = e.GetPriceAt(r.SymbolCode, from)
		results[i].Last, _ = e.GetLastPrice(r.SymbolCode, &status, nil)
		results[i].PrevVolume, _ = e.GetVolume(r.SymbolCode, &status, &from, &to)
	}
	return results, nil
}

func (e *Exchange) GetPriceAt(symbol string, from time.Time) (uint64, error) {
	var result struct {
		Price uint64
	}
	if err := e.db.Table("exchange_orders").Joins("JOIN exchange_markets em ON em.id = exchange_orders.market_id").Where("exchange_orders.updated_at >= ?", from).Where("em.symbol_code = ?", symbol).Where("status = ?", int(models.Filled)).Order("exchange_orders.updated_at ASC").Limit(1).Select("price").Scan(&result).Error; err != nil {
		return 0, errors.Wrap(err, "e.db.Table")
	}
	return result.Price, nil
}

func (e *Exchange) GetVolume(symbol string, status *models.OrderStatus, from, to *time.Time) (uint64, error) {
	var result struct {
		Volume uint64
	}
	query := e.db.Table("exchange_orders").Joins("JOIN exchange_markets em ON em.id = exchange_orders.market_id").Where("symbol_code = ?", symbol)
	if from != nil {
		query = query.Where("exchange_orders.updated_at >= ?", *from)
	}
	if to != nil {
		query = query.Where("exchange_orders.updated_at <= ?", *to)
	}
	if status != nil {
		query = query.Where("status = ?", int(*status))
	}
	query = query.Select("sum(quantity) as volume")

	if err := query.Scan(&result).Error; err != nil {
		return 0, errors.Wrap(err, "e.db.Table")
	}
	return result.Volume, nil
}

func (e *Exchange) GetLastPrice(symbol string, status *models.OrderStatus, side *models.OrderSide) (uint64, error) {
	var result struct {
		Price uint64
	}
	query := e.db.Table("exchange_orders").Joins("JOIN exchange_markets em ON em.id = exchange_orders.market_id").Where("em.symbol_code = ?", symbol)
	if status != nil {
		query = query.Where("status = ?", int(*status))
	}
	if side != nil {
		query = query.Where("side = ?", int(*side))
	}
	query = query.Order("exchange_orders.updated_at DESC").Limit(1).Select("price")

	if err := query.Scan(&result).Error; err != nil {
		return 0, errors.Wrap(err, "e.db.Table")
	}
	return result.Price, nil
}

func (e *Exchange) GetHighLowPrice(symbol string, status *models.OrderStatus, side *models.OrderSide, from, to *time.Time) (uint64, uint64, error) {
	var result struct {
		High, Low uint64
	}

	query := e.db.Table("exchange_orders").Joins("JOIN exchange_markets em ON em.id = exchange_orders.market_id").Where("em.symbol_code = ?", symbol)
	if from != nil {
		query = query.Where("exchange_orders.updated_at >= ?", *from)
	}
	if to != nil {
		query = query.Where("exchange_orders.updated_at <= ?", *to)
	}
	if status != nil {
		query = query.Where("exchange_orders.status = ?", int(*status))
	}
	if side != nil {
		query = query.Where("exchange_orders.side = ?", int(*side))
	}
	query = query.Select("max(price) as high, min(price) as low")

	if err := query.Scan(&result).Error; err != nil {
		return 0, 0, errors.Wrap(err, "query.Scan")
	}
	return result.High, result.Low, nil
}

func (e *Exchange) MarketRates() ([]*MarketRate, error) {
	markets, err := e.ListMarkets("")
	if err != nil {
		return nil, errors.Wrap(err, "e.ListMarkets")
	}

	rates := make([]*MarketRate, 0, len(markets))
	for _, m := range markets {
		rate, err := e.MarketRatesBySymbol(m.SymbolCode)
		if err != nil {
			return nil, errors.Wrapf(err, "e.MarketRatesBySymbol %s", m.SymbolCode)
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

func (e *Exchange) MarketRatesBySymbol(symbol string) (*MarketRate, error) {
	status := models.Filled
	last, err := e.GetLastPrice(symbol, &status, nil)
	if err != nil {
		return nil, errors.Wrap(err, "e.GetLastPrice")
	}

	status = models.New
	side := models.Buy
	bid, err := e.GetLastPrice(symbol, &status, &side)
	if err != nil {
		return nil, errors.Wrap(err, "e.GetLastPrice")
	}

	side = models.Sell
	ask, err := e.GetLastPrice(symbol, &status, &side)
	if err != nil {
		return nil, errors.Wrap(err, "e.GetLastPrice")
	}

	volume, err := e.GetVolume(symbol, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "e.GetVolume")
	}

	from, to := time.Now().Add(-24*time.Hour), time.Now()
	high, low, err := e.GetHighLowPrice(symbol, nil, nil, &from, &to)
	if err != nil {
		return nil, errors.Wrap(err, "e.GetHighLow")
	}
	return &MarketRate{
		SymbolCode: symbol,
		Last:       last,
		Bid:        bid,
		Ask:        ask,
		Volume:     volume,
		High24h:    high,
		Low24h:     low,
	}, nil
}

type MarketRate struct {
	SymbolCode      string
	Last            uint64
	Bid, Ask        uint64
	Volume          uint64
	High24h, Low24h uint64
}

type SymbolRate struct {
	SymbolCode string
	Volume     uint64
	Last       uint64
	High, Low  uint64
	PrevPrice  uint64
	PrevVolume uint64
}
