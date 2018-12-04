package voting

import (
	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/pkg/errors"
)

func (p *VotingDao) CreateVotingProposalDCB(b *models.VotingProposalDCB) (*models.VotingProposalDCB, error) {
	if err := p.db.Create(b).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Create")
	}
	return b, nil
}

func (p *VotingDao) CreateVotingProposalGOV(b *models.VotingProposalGOV) (*models.VotingProposalGOV, error) {
	if err := p.db.Create(b).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Create")
	}
	return b, nil
}

func (p *VotingDao) CreateVotingProposalGOVVote(b *models.VotingProposalGOVVote) (*models.VotingProposalGOVVote, error) {
	if err := p.db.Create(b).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Create")
	}
	return b, nil
}

func (p *VotingDao) CreateVotingProposalDCBVote(b *models.VotingProposalDCBVote) (*models.VotingProposalDCBVote, error) {
	if err := p.db.Create(b).Error; err != nil {
		return nil, errors.Wrap(err, "b.db.Create")
	}
	return b, nil
}

func (p *VotingDao) GetDCBProposal(id int) (*models.VotingProposalDCB, error) {
	var v models.VotingProposalDCB
	if err := p.db.Preload("User").Preload("VotingProposalDCBVotes").Preload("VotingProposalDCBVotes.Voter").First(&v, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "p.db.Preload")
	}
	return &v, nil
}

func (p *VotingDao) GetGOVProposal(id int) (*models.VotingProposalGOV, error) {
	var v models.VotingProposalGOV
	if err := p.db.Preload("User").Preload("VotingProposalGOVVotes").Preload("VotingProposalGOVVotes.Voter").First(&v, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "p.db.Preload")
	}
	return &v, nil
}

func (p *VotingDao) GetDCBProposals(limit, page *int) ([]*models.VotingProposalDCB, error) {
	var vs []*models.VotingProposalDCB

	query := p.db.Preload("User").Preload("VotingProposalDCBVotes").Preload("VotingProposalDCBVotes.Voter")
	if limit != nil && page != nil {
		offset := (*page)*(*limit) - *limit
		query = query.Limit(*limit).Offset(offset)
	}
	query = query.Order("created_at DESC")

	if err := query.Find(&vs).Error; err != nil {
		return nil, errors.Wrap(err, "e.db.Where")
	}
	return vs, nil
}

func (p *VotingDao) GetGOVProposals(limit, page *int) ([]*models.VotingProposalGOV, error) {
	var vs []*models.VotingProposalGOV

	query := p.db.Preload("User").Preload("VotingProposalGOVVotes").Preload("VotingProposalGOVVotes.Voter")
	if limit != nil && page != nil {
		offset := (*page)*(*limit) - *limit
		query = query.Limit(*limit).Offset(offset)
	}
	query = query.Order("created_at DESC")

	if err := query.Find(&vs).Error; err != nil {
		return nil, errors.Wrap(err, "e.db.Where")
	}
	return vs, nil
}

func (p *VotingDao) GetProposalDCBVote(id uint) (int, error) {
	var result struct {
		TotalVote int
	}
	if err := p.db.Model(&models.VotingProposalDCBVote{}).Where("voting_proposal_dcb_id = ?", id).Select("count(*) as total_vote").Group("voting_proposal_dcb_id").Scan(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, errors.Wrap(err, "p.db.Model")
	}
	return result.TotalVote, nil
}

func (p *VotingDao) GetProposalGOVVote(id uint) (int, error) {
	var result struct {
		TotalVote int
	}
	if err := p.db.Model(&models.VotingProposalGOVVote{}).Where("voting_proposal_gov_id = ?", id).Select("count(*) as total_vote").Group("voting_proposal_gov_id").Scan(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, errors.Wrap(err, "p.db.Model")
	}
	return result.TotalVote, nil
}
