package reserve

import "github.com/jinzhu/gorm"

type ReserveDao struct {
	db *gorm.DB
}

func NewReserveDao(db *gorm.DB) *ReserveDao {
	return &ReserveDao{db}
}
