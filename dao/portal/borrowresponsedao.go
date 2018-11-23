package portal

import (
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/pkg/errors"
)

func (p *PortalDao) CreateBorrowResponse(b *models.BorrowResponse) (*models.BorrowResponse, error) {
	if err := p.db.Create(b).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Create")
	}
	return b, nil
}

func (p *PortalDao) ListAllBorrowsResponse(portal_borrow_id int) ([]*models.BorrowResponse, error) {
	var (
		bs []*models.BorrowResponse
	)

	query := p.db
	query = query.Where("portal_borrow_id = ?", portal_borrow_id)
	query = query.Find(&bs)

	if err := query.Find(&bs).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Where.Find")
	}
	return bs, nil
}
