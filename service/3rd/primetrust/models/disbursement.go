package models

import "github.com/mongodb/mongo-go-driver/bson/decimal"

type DisbursementAttributes struct {
	AccountID         string                  `json:"account-id"`
	Amount            decimal.Decimal128      `json:"amount"`
	CustomerReference string                  `json:"customer-reference"`
	Description       string                  `json:"description"`
	PaymentMethodID   string                  `json:"payment-method-id"`
	PaymentMethod     PaymentMethodAttributes `json:"payment-method"`
	ID                string                  `json:"id"`
	SpecialType       string                  `json:"contact-email"`
	Status            string                  `json:"status"`
	ContactEmail      string                  `json:"contact-email"`
	ContactName       string                  `json:"contact-name"`
}

type DisbursementData struct {
	ID            string                 `json:"id,omitempty"`
	Type          string                 `json:"type,omitempty"`
	Attributes    DisbursementAttributes `json:"attributes"`
	Links         Links                  `json:"links"`
	Relationships Relationships          `json:"relationships"`
}

type Disbursement struct {
	Data DisbursementData `json:"data"`
}

type DisbursementResponse struct {
	CollectionResponse
	Data []DisbursementData `json:"data"`
}
