package reserve

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

func (r *Reserve) CreateReserveDisbursementRequest(rdr *models.ReserveDisbursementRequest) (*models.ReserveDisbursementRequest, error) {
	if err := r.db.Create(rdr).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Create")
	}
	return rdr, nil
}

func (r *Reserve) UpdateReserveDisbursementRequest(rdr *models.ReserveDisbursementRequest) (*models.ReserveDisbursementRequest, error) {
	if err := r.db.Save(rdr).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Update")
	}
	return rdr, nil
}

func (r *Reserve) DeleteReserveDisbursementRequest(rdr *models.ReserveDisbursementRequest) error {
	if err := r.db.Delete(rdr).Error; err != nil {
		return errors.Wrap(err, "r.db.Delete")
	}
	return nil
}

func (r *Reserve) FindReserveDisbursementRequestByID(id int) (*models.ReserveDisbursementRequest, error) {
	var rdr models.ReserveDisbursementRequest
	if err := r.db.First(&rdr, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "r.db.First")
	}
	return &rdr, nil
}

func (r *Reserve) FindAllReserveDisbursementRequest(filter *map[string]interface{}, page, limit int) ([]*models.ReserveDisbursementRequest, error) {
	var (
		rdrs   []*models.ReserveDisbursementRequest
		offset = page*limit - limit
	)

	query := r.db.Table("reserve_disbursement_request")
	query = query.Limit(limit).Offset(offset)
	if filter != nil {
		query = query.Where(filter)
	}

	if err := query.Find(&rdrs).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Find")
	}
	return rdrs, nil
}
