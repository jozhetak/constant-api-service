package reserve

import "github.com/jinzhu/gorm"

type Reserve struct {
	db *gorm.DB
}

func NewReserve(db *gorm.DB) *ReserveDao {
	return &Reserve{db}
}
