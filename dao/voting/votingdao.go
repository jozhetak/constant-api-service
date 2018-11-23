package voting

import (
	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/pkg/errors"
)

type VotingDao struct {
	db *gorm.DB
}

func NewVoting(db *gorm.DB) *VotingDao {
	return &VotingDao{db}
}

func (p *VotingDao) CreateVotingBoardCandidate(b *models.VotingBoardCandidate) (*models.VotingBoardCandidate, error) {
	if err := p.db.Create(b).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Create")
	}
	return b, nil
}

func (p *VotingDao) UpdateVotingBoardCandidate(b *models.VotingBoardCandidate) (*models.VotingBoardCandidate, error) {
	if err := p.db.Save(b).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Update")
	}
	return b, nil
}

func (p *VotingDao) DeleteVotingBoardCandidate(b *models.VotingBoardCandidate) (error) {
	if err := p.db.Delete(&b).Error; err != nil {
		return errors.Wrap(err, "b.db.Delete")
	}
	return nil
}

func (p *VotingDao) FindBorrowByID(id int) (*models.VotingBoardCandidate, error) {
	var b models.VotingBoardCandidate
	if err := p.db.First(&b, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "b.db.First")
	}
	return &b, nil
}