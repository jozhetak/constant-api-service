package reserve

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

func (r *Reserve) CreateReserveContributionRequest(rcr *models.ReserveContributionRequest) (*models.ReserveContributionRequest, error) {
	if err := r.db.Create(rcr).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Create")
	}
	return rcr, nil
}

func (r *Reserve) UpdateReserveContributionRequest(rcr *models.ReserveContributionRequest) (*models.ReserveContributionRequest, error) {
	if err := r.db.Save(rcr).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Update")
	}
	return rcr, nil
}

func (r *Reserve) DeleteReserveContributionRequest(rcr *models.ReserveContributionRequest) error {
	if err := r.db.Delete(rcr).Error; err != nil {
		return errors.Wrap(err, "r.db.Delete")
	}
	return nil
}

func (r *Reserve) FindReserveContributionRequestByID(id int) (*models.ReserveContributionRequest, error) {
	var rcr models.ReserveContributionRequest
	if err := r.db.First(&rcr, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "r.db.First")
	}
	return &rcr, nil
}

func (r *Reserve) FindAllReserveContributionRequest(filter *map[string]interface{}, page int, limit int) ([]*models.ReserveContributionRequest, error) {
	var (
		rcrs   []*models.ReserveContributionRequest
		offset = page*limit - limit
	)

	query := r.db.Table("reserve_contribution_request")
	query = query.Limit(limit).Offset(offset)
	if filter != nil {
		query = query.Where(*filter)
	}

	if err := query.Find(&rcrs).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Find")
	}
	return rcrs, nil
}
