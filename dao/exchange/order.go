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

func (e *Exchange) UserOrderHistory(u *models.User, symbol string, status *models.OrderStatus, limit, page int) ([]*models.Order, error) {
	var (
		orders []*models.Order
		offset = page*limit - limit
	)
	query := e.db.Preload("Market", "symbol_code = ?", symbol).Where("user_id = ?", u.ID)
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	query = query.Order("created_at DESC").Limit(limit).Offset(offset)

	if err := query.Find(&orders).Error; err != nil {
		return nil, errors.Wrap(err, "e.db.Where")
	}
	return orders, nil
}
