package voting

import (
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
