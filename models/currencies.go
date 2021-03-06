package models

import "github.com/jinzhu/gorm"

type Currency struct {
	gorm.Model
	Name        string
	TokenID     string
	TokenName   string
	TokenSymbol string
}

func (*Currency) TableName() string {
	return "exchange_currencies"
}
