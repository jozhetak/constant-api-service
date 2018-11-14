package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserType int

const (
	InvalidUserType UserType = iota
	Borrower
	Lender
)

func (u UserType) String() string {
	return [...]string{"borrower", "lender"}[u]
}

type User struct {
	gorm.Model
	FirstName      string
	LastName       string
	UserName       string
	Email          string
	Password       string
	PaymentAddress string
	ReadonlyKey    string
	PrivKey        string
	Type           UserType
}

func GetUserType(kind string) UserType {
	switch kind {
	case "borrower":
		return Borrower
	case "lender":
		return Lender
	default:
		return InvalidUserType
	}
}

type UserRecovery struct {
	gorm.Model

	User   *User
	UserID int

	Token     string
	IsValid   bool
	ExpiredAt time.Time
}
