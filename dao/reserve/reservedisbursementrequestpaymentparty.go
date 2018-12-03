package reserve

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

func (r *Reserve) CreateReserveDisbursementRequestPaymentParty(rdrpp *models.ReserveDisbursementRequestPaymentParty) (*models.ReserveDisbursementRequestPaymentParty, error) {
	if err := r.db.Create(rdrpp).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Create")
	}
	return rdrpp, nil
}

func (r *Reserve) UpdateReserveDisbursementRequestPaymentParty(rdrpp *models.ReserveDisbursementRequestPaymentParty) (*models.ReserveDisbursementRequestPaymentParty, error) {
	if err := r.db.Save(rdrpp).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Update")
	}
	return rdrpp, nil
}

func (r *Reserve) DeleteReserveDisbursementRequestPaymentParty(rdrpp *models.ReserveDisbursementRequestPaymentParty) error {
	if err := r.db.Delete(rdrpp).Error; err != nil {
		return errors.Wrap(err, "r.db.Delete")
	}
	return nil
}

func (r *Reserve) FindReserveDisbursementRequestPaymentPartyByID(id int) (*models.ReserveDisbursementRequestPaymentParty, error) {
	var rdrpp models.ReserveDisbursementRequestPaymentParty
	if err := r.db.First(&rdrpp, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "r.db.First")
	}
	return &rdrpp, nil
}

func (r *Reserve) FindAllReserveDisbursementRequestPaymentParty(filter *map[string]interface{}) ([]*models.ReserveDisbursementRequestPaymentParty, error) {
	var rdrpps []*models.ReserveDisbursementRequestPaymentParty

	query := r.db
	if filter != nil {
		query = query.Where(filter)
	}

	if err := query.Find(&rdrpps).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Find")
	}
	return rdrpps, nil
}
