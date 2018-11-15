package exchange

import (
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/pkg/errors"
)

func (e *Exchange) CreateOrder(o *models.Order) (*models.Order, error) {
	if err := e.db.Create(o).Error; err != nil {
		return nil, errors.Wrap(err, "e.db.Create")
	}
	return o, nil
}
