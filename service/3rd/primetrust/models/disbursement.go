package models

import "time"

type DisbursementAttributes struct {
	AccountID         string        `json:"account-id"`
	Amount            float64       `json:"amount"`
	ContactEmail      string        `json:"contact-email"`
	ContactName       string        `json:"contact-name"`
	CreatedAt         *time.Time    `json:"created-at"`
	PaymentDetails    *string       `json:"payment-details"`
	ReferenceNumber   string        `json:"reference-number"`
	Status            string        `json:"status"`
	TransactionNumber *string       `json:"transaction-number"`
	PaymentMethodID   string        `json:"payment-method-id"`
	Links             Links         `json:"links"`
	Relationships     Relationships `json:"relationships"`
}

type DisbursementData struct {
	ID         string                 `json:"id,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Attributes DisbursementAttributes `json:"attributes"`
}

type Disbursement struct {
	Data ContributionData `json:"data"`
}
