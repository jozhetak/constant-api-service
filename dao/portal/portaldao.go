package portal

import (
	"github.com/jinzhu/gorm"
)

type PortalDao struct {
	db *gorm.DB
}

func NewPortal(db *gorm.DB) *PortalDao {
	return &PortalDao{db}
}
