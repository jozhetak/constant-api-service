package reserve

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

func (r *Reserve) CreateReserveContributionRequestPaymentParty(rcrpp *models.ReserveContributionRequestPaymentParty) (*models.ReserveContributionRequestPaymentParty, error) {
	if err := r.db.Create(rcrpp).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Create")
	}
	return rcrpp, nil
}

func (r *Reserve) UpdateReserveContributionRequestPaymentParty(rcrpp *models.ReserveContributionRequestPaymentParty) (*models.ReserveContributionRequestPaymentParty, error) {
	if err := r.db.Save(rcrpp).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Update")
	}
	return rcrpp, nil
}

func (r *Reserve) DeleteReserveContributionRequestPaymentParty(rcrpp *models.ReserveContributionRequestPaymentParty) error {
	if err := r.db.Delete(rcrpp).Error; err != nil {
		return errors.Wrap(err, "r.db.Delete")
	}
	return nil
}

func (r *Reserve) FindReserveContributionRequestPaymentPartyByID(id int) (*models.ReserveContributionRequestPaymentParty, error) {
	var rcrpp models.ReserveContributionRequestPaymentParty
	if err := r.db.First(&rcrpp, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "r.db.First")
	}
	return &rcrpp, nil
}

func (r *Reserve) FindAllReserveContributionRequestPaymentParty(filter *map[string]interface{}) ([]*models.ReserveContributionRequestPaymentParty, error) {
	var rcrpps []*models.ReserveContributionRequestPaymentParty

	query := r.db
	if filter != nil {
		query = query.Where(filter)
	}

	if err := query.Find(&rcrpps).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Find")
	}
	return rcrpps, nil
}
