package portal

import (
	"github.com/jinzhu/gorm"
)

type Portal struct {
	db *gorm.DB
}

func NewPortal(db *gorm.DB) *Portal {
	return &Portal{db}
}
