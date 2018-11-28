package exchange

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

func (e *Exchange) CreateOrder(o *models.Order) (*models.Order, error) {
	if err := e.db.Create(o).Error; err != nil {
		return nil, errors.Wrap(err, "e.db.Create")
	}
	return o, nil
}

func (e *Exchange) OrderHistory(symbol string, status *models.OrderStatus, side *models.OrderSide, limit, page *int, u *models.User) ([]*models.Order, error) {
	var orders []*models.Order

	query := e.db.Preload("Market").Table("exchange_orders").Joins("JOIN exchange_markets em ON em.id = exchange_orders.market_id").Where("em.symbol_code = ?", symbol)
	if u != nil {
		query = query.Where("user_id = ?", u.ID)
	}
	if status != nil {
		query = query.Where("status = ?", int(*status))
	}
	if limit != nil && page != nil {
		offset := (*page)*(*limit) - *limit
		query = query.Limit(*limit).Offset(offset)
	}
	query = query.Order("created_at DESC")

	if err := query.Find(&orders).Error; err != nil {
		return nil, errors.Wrap(err, "e.db.Where")
	}
	return orders, nil
}

func (e *Exchange) FindOrderByID(id int) (*models.Order, error) {
	var o models.Order
	if err := e.db.Preload("User").Preload("Market").Preload("Market.BaseCurrency").Preload("Market.QuoteCurrency").First(&o, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "b.db.First")
	}
	return &o, nil
}

func (e *Exchange) SetFilledOrders(orders ...*models.Order) (err error) {
	tx := e.db.Begin()
	if tErr := tx.Error; tErr != nil {
		err = errors.Wrap(tErr, "tx.Error")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = errors.Wrap(tx.Commit().Error, "tx.Commit")
	}()

	for _, o := range orders {
		if tErr := tx.Model(o).Update("status", int(models.Filled)).Error; tErr != nil {
			err = errors.Wrap(tErr, "tx.Model")
			return
		}
	}
	return
}

func (e *Exchange) FindOrdersInMarkets(markets []*models.Market, status *models.OrderStatus, side *models.OrderSide) ([]*models.Order, error) {
	var orders []*models.Order

	marketIDs := make([]uint, 0, len(markets))
	for _, m := range markets {
		marketIDs = append(marketIDs, m.ID)
	}
	query := e.db.Preload("Market").Table("exchange_orders").Where("market_id IN (?)", marketIDs)

	if status != nil {
		query = query.Where("status = ?", int(*status))
	}
	if side != nil {
		query = query.Where("side = ?", int(*side))
	}
	query = query.Order("created_at DESC")

	if err := query.Find(&orders).Error; err != nil {
		return nil, errors.Wrap(err, "e.db.Where")
	}
	return orders, nil
}
