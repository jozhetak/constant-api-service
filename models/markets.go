package models

import "github.com/jinzhu/gorm"

type MarketState int

const (
	InvalidMarketState MarketState = iota
	Online
)

func (m MarketState) String() string {
	return [...]string{"invalid", "online"}[m]
}

type Market struct {
	gorm.Model

	// BaseCurrency string
	BaseCurrencyID int
	BaseCurrency   *Currency `gorm:"foreignkey:BaseCurrencyID"`

	// QuoteCurrency   string
	QuoteCurrencyID int
	QuoteCurrency   *Currency `gorm:"foreignkey:QuoteCurrencyID"`

	DisplayName          string
	State                MarketState
	SymbolCode           string
	Icon                 string
	TradeEnabled         bool
	FeePrecision         int
	TradePricePrecision  int
	TradeTotalPrecision  int
	TradeAmountPrecision int
}

func (*Market) TableName() string {
	return "exchange_markets"
}
