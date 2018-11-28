package models

import (
	"strings"
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
	return [...]string{"invalid", "buy", "sell"}[o]
}

type OrderStatus int

const (
	InvalidOrderStatus OrderStatus = iota
	New
	Filled
)

func (o OrderStatus) String() string {
	return [...]string{"invalid", "new", "filled"}[o]
}

type OrderType int

const (
	InvalidOrderType OrderType = iota
	Limit
	MarketType
)

func (o OrderType) String() string {
	return [...]string{"invalid", "limit", "market"}[o]
}

type Order struct {
	gorm.Model

	User   *User
	UserID int

	Market   *Market
	MarketID int

	Price    uint64
	Quantity uint64
	Type     OrderType
	Status   OrderStatus
	Side     OrderSide
	Time     time.Time
}

func (*Order) TableName() string {
	return "exchange_orders"
}

func GetOrderType(t string) OrderType {
	switch strings.ToLower(t) {
	case "limit":
		return Limit
	case "market":
		return MarketType
	default:
		return InvalidOrderType
	}
}

func GetOrderSide(s string) OrderSide {
	switch strings.ToLower(s) {
	case "buy":
		return Buy
	case "sell":
		return Sell
	default:
		return InvalidOrderSide
	}
}

func GetOrderStatus(s string) OrderStatus {
	switch strings.ToLower(s) {
	case "new":
		return New
	case "filled":
		return Filled
	default:
		return InvalidOrderStatus
	}
}
