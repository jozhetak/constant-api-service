package portal

import (
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func (p *PortalDao) CreateBorrow(b *models.Borrow) (*models.Borrow, error) {
	if err := p.db.Create(b).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Create")
	}
	return b, nil
}

func (p *PortalDao) UpdateBorrow(b *models.Borrow) (*models.Borrow, error) {
	if err := p.db.Save(b).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Update")
	}
	return b, nil
}

func (p *PortalDao) DeleteBorrow(b *models.Borrow) (error) {
	if err := p.db.Delete(&b).Error; err != nil {
		return errors.Wrap(err, "b.db.Delete")
	}
	return nil
}

func (p *PortalDao) ListBorrowByUser(paymentAddress string, state *models.BorrowStatus, limit, page int) ([]*models.Borrow, error) {
	var (
		bs     []*models.Borrow
		offset = page*limit - limit
	)

	query := p.db.Where("payment_address = ?", paymentAddress)
	query = query.Preload("BorrowResponses").Preload("User")
	if state != nil {
		query = query.Where("state = ?", *state)
	}
	query = query.Order("created_at DESC").Limit(limit).Offset(offset)

	if err := query.Find(&bs).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Where.Find")
	}
	return bs, nil
}

func (p *PortalDao) ListAllBorrows(state *models.BorrowStatus, limit, page int) ([]*models.Borrow, error) {
	var (
		bs     []*models.Borrow
		offset = page*limit - limit
	)

	query := p.db.Limit(limit).Offset(offset)
	if state != nil {
		query = query.Preload("portal_borrows.BorrowResponse").Where("state = ?", *state)
	}
	query = query.Find(&bs)

	if err := query.Find(&bs).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Where.Find")
	}
	return bs, nil
}

func (p *PortalDao) FindBorrowByID(id int) (*models.Borrow, error) {
	var b models.Borrow
	if err := p.db.First(&b, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "b.db.First")
	}
	return &b, nil
}
