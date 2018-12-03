package exchange

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

func (r *Reserve) CreateReserveDisbursementRequest(rcr *models.ReserveDisbursementRequest) (*models.ReserveDisbursementRequest, error) {
	if err := r.db.Create(rcr).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Create")
	}
	return rcr, nil
}

func (r *Reserve) UpdateReserveDisbursementRequest(rcr *models.ReserveDisbursementRequest) (*models.ReserveDisbursementRequest, error) {
	if err := r.db.Save(crc).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Update")
	}
	return rcr, nil
}

func (r *Reserve) DeleteReserveDisbursementRequest(rcr *models.ReserveDisbursementRequest) error {
	if err := r.db.Delete(rcr).Error; err != nil {
		return errors.Wrap(err, "r.db.Delete")
	}
	return nil
}

func (r *Reserve) FindReserveDisbursementRequestByID(id int) (*models.ReserveDisbursementRequest, error) {
	var rcr models.ReserveDisbursementRequest
	if err := r.db.First(&b, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "r.db.First")
	}
	return &b, nil
}

func (r *Reserve) FindAllReserveDisbursementRequest(filter *map[string]interface{}) ([]*models.ReserveDisbursementRequest, error) {
	var rcrs []*models.ReserveDisbursementRequest

	query := r.db.Table("reserve_disbursement_request")
	if filter != nil {
		query = query.Where(filter)
	}

	if err := query.Find(&rcrs).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Find")
	}
	return rcrs, nil
}
