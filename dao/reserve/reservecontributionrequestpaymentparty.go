package exchange

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

func (r *Reserve) CreateReserveContributionRequestPaymentParty(rcr *models.ReserveContributionRequestPaymentParty) (*models.ReserveContributionRequestPaymentParty, error) {
	if err := r.db.Create(rcr).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Create")
	}
	return rcr, nil
}

func (r *Reserve) UpdateReserveContributionRequestPaymentParty(rcr *models.ReserveContributionRequestPaymentParty) (*models.ReserveContributionRequestPaymentParty, error) {
	if err := r.db.Save(crc).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Update")
	}
	return rcr, nil
}

func (r *Reserve) DeleteReserveContributionRequestPaymentParty(rcr *models.ReserveContributionRequestPaymentParty) error {
	if err := r.db.Delete(rcr).Error; err != nil {
		return errors.Wrap(err, "r.db.Delete")
	}
	return nil
}

func (r *Reserve) FindReserveContributionRequestPaymentPartyByID(id int) (*models.ReserveContributionRequestPaymentParty, error) {
	var rcr models.ReserveContributionRequestPaymentParty
	if err := r.db.First(&b, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "r.db.First")
	}
	return &b, nil
}

func (r *Reserve) FindAllReserveContributionRequestPaymentParty(filter *map[string]interface{}) ([]*models.ReserveContributionRequestPaymentParty, error) {
	var rcrs []*models.ReserveContributionRequestPaymentParty

	query := r.db
	if filter != nil {
		query = query.Where(filter)
	}

	if err := query.Find(&rcrs).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Find")
	}
	return rcrs, nil
}
