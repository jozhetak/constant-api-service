package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/conf"
)

func Init(config *config.Config) (*gorm.DB, error) {
	databaseConn, err := gorm.Open("mysql", config.Db)
	databaseConn.LogMode(true)
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Open")
	}
	// skip save associations of gorm -> manual save by code
	databaseConn = databaseConn.Set("gorm:save_associations", false)
	databaseConn.DB().SetMaxOpenConns(20)
	databaseConn.DB().SetMaxIdleConns(10)
	return databaseConn, err
}
