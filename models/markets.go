package models

import "github.com/jinzhu/gorm"

type Market struct {
	gorm.Model
	BaseCurrency   string
	MarketCurrency string
	Symbol         string
}

func (*Market) TableName() string {
	return "exchange_markets"
}
