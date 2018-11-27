package models

type PaymentMethod struct {
	AchCheckType      string `json:"ach-check-type"`
	BankAccountName   string `json:"bank-account-name"`
	BankAccountType   string `json:"bank-account-type"`
	BankAccountNumber string `json:"bank-account-number"`
	ContactEmail      string `json:"contact-email"`
	ContactName       string `json:"contact-name"`
	Inactive          bool   `json:"inactive"`
	Last4             string `json:"last-4"`
	PaymentType       string `json:"payment-type"`
	RoutingNumber     string `json:"routing-number"`
}
