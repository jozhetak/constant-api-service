package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OrderSide int

const (
	InvalidOrderSide OrderSide = iota
	Buy
	Sell
)

func (o OrderSide) String() string {
	return [...]string{"INVALID", "BUY", "SELL"}[o]
}

type OrderStatus int

const (
	InvalidOrderStatus OrderStatus = iota
	New
	Filled
)

func (o OrderStatus) String() string {
	return [...]string{"INVALID", "NEW", "FILLED"}[o]
}

type OrderType int

const (
	InvalidOrderType OrderType = iota
	Limit
)

func (o OrderType) String() string {
	return [...]string{"INVALID", "LIMIT"}[o]
}

type Order struct {
	gorm.Model

	User   *User
	UserID int

	Market   *Market
	MarketID int

	Price    float64
	Quantity uint
	Type     OrderType
	Status   OrderStatus
	Side     OrderSide
	Time     time.Time
}

func (*Order) TableName() string {
	return "exchange_orders"
}

func GetOrderType(t string) OrderType {
	switch t {
	case "limit":
		return Limit
	default:
		return InvalidOrderType
	}
}

func GetOrderSide(s string) OrderSide {
	switch s {
	case "buy":
		return Buy
	case "sell":
		return Sell
	default:
		return InvalidOrderSide
	}
}
