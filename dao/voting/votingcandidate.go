package voting

import (
	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/pkg/errors"
)

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

func (p *VotingDao) DeleteVotingBoardCandidate(b *models.VotingBoardCandidate) error {
	if err := p.db.Delete(&b).Error; err != nil {
		return errors.Wrap(err, "b.db.Delete")
	}
	return nil
}

func (p *VotingDao) FindVotingBoardCandidateByID(id int) (*models.VotingBoardCandidate, error) {
	var b models.VotingBoardCandidate
	if err := p.db.First(&b, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "b.db.First")
	}
	return &b, nil
}

func (p *VotingDao) FindVotingBoardCandidateByUser(user models.User) (*models.VotingBoardCandidate, error) {
	var b models.VotingBoardCandidate
	if err := p.db.Preload("User").Where("user_id = ?", user.ID).First(&b).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "b.db.First")
	}
	return &b, nil
}

type VotingCandidateFilter struct {
	BoardType      int
	PaymentAddress string
}

func (p *VotingDao) Filter(filter *VotingCandidateFilter) ([]*models.VotingBoardCandidate, error) {
	var b []*models.VotingBoardCandidate
	query := p.db.Preload("User")

	switch models.BoardCandidateType(filter.BoardType) {
	case models.DCB:
		if filter.PaymentAddress == "" {
			query = query.Where("dcb IS NOT NULL")
		} else {
			query = query.Where("dcb = ?", filter.PaymentAddress)
		}
	case models.GOV:
		if filter.PaymentAddress == "" {
			query = query.Where("gov IS NOT NULL")
		} else {
			query = query.Where("gov = ?", filter.PaymentAddress)
		}
	case models.CMB:
		if filter.PaymentAddress == "" {
			query = query.Where("cmb IS NOT NULL")
		} else {
			query = query.Where("cmb = ?", filter.PaymentAddress)
		}
	}
	if err := query.Find(&b).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "b.db.First")
	}
	return b, nil
}

func (p *VotingDao) CreateVotingBoardVote(b *models.VotingBoardVote) (*models.VotingBoardVote, error) {
	if err := p.db.Create(b).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Create")
	}
	return b, nil
}

func (p *VotingDao) GetTotalVote(id uint) (int, error) {
	var result struct {
		TotalVote int
	}
	if err := p.db.Model(&models.VotingBoardVote{}).Where("voting_board_candidate_id = ?", id).Select("count(*) as total_vote").Group("voting_board_candidate_id").Scan(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, errors.Wrap(err, "p.db.Model")
	}
	return result.TotalVote, nil
}

func (p *VotingDao) FindCandidateByID(id int) (*models.VotingBoardCandidate, error) {
	var c models.VotingBoardCandidate
	if err := p.db.First(&c, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "p.db.First")
	}
	return &c, nil
}
