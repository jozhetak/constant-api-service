package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/constant-api-service/conf"
)

func Init(config *config.Config) (*gorm.DB, error) {
	var databaseConn *gorm.DB
	//open a db connection
	databaseConn, err := gorm.Open("mysql", config.Db)
	databaseConn.LogMode(true)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// skip save associations of gorm -> manual save by code
	databaseConn = databaseConn.Set("gorm:save_associations", false)
	databaseConn.DB().SetMaxOpenConns(20)
	databaseConn.DB().SetMaxIdleConns(10)
	return databaseConn, err
}
