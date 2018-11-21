package exchange

import (
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

	query := e.db.Table("exchange_orders").Joins("JOIN exchange_markets em ON em.id = exchange_orders.market_id").Where("em.symbol_code = ?", symbol)
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
