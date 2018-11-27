package voting

import (
	"github.com/jinzhu/gorm"
)

type VotingDao struct {
	db *gorm.DB
}

func NewVoting(db *gorm.DB) *VotingDao {
	return &VotingDao{db}
}
