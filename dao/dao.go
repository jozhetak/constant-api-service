package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/models"
)

var (
	userTables     = []interface{}{(*models.User)(nil), (*models.UserVerification)(nil)}
	portalTables   = []interface{}{(*models.Borrow)(nil)}
	exchangeTables = []interface{}{(*models.Currency)(nil), (*models.Market)(nil), (*models.Order)(nil)}
	votingTables   = []interface{}{(*models.VotingBoardCandidate)(nil), (*models.VotingBoardVote)(nil), (*models.VotingProposalDCB)(nil), (*models.VotingProposalDCBVote)(nil), (*models.VotingProposalGOV)(nil), (*models.VotingProposalGOVVote)(nil)}
)

func AutoMigrate(db *gorm.DB) error {
	allTables := make([]interface{}, 0, len(userTables)+len(exchangeTables)+len(votingTables))
	allTables = append(allTables, userTables...)
	allTables = append(allTables, portalTables...)
	allTables = append(allTables, exchangeTables...)
	allTables = append(allTables, votingTables...)
	if err := db.AutoMigrate(allTables...).Error; err != nil {
		return errors.Wrap(err, "db.AutoMigrate")
	}
	return nil
}
