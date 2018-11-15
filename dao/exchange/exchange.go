package exchange

import (
	"github.com/jinzhu/gorm"
)

type Exchange struct {
	db *gorm.DB
}

func NewExchange(db *gorm.DB) *Exchange {
	return &Exchange{db}
}
