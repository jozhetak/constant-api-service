package models

import "time"

type ContributionAttributes struct {
	AccountID           string        `json:"account-id"`
	Amount              float64       `json:"amount"`
	ContactID           string        `json:"contact-id"`
	ContactEmail        string        `json:"contact-email"`
	ContactName         string        `json:"contact-name"`
	PaymentType         string        `json:"payment-type"`
	AmountExpected      float64       `json:"amount-expected"`
	CreatedAt           time.Time     `json:"created-at"`
	Message             *string       `json:"message"`
	PaymentDetails      *string       `json:"payment-details"`
	ReferenceNumber     string        `json:"reference-number"`
	SpecialInstructions *string       `json:"special-instructions"`
	SpecialType         *string       `json:"special-type"`
	Status              string        `json:"status"`
	TransactionNumber   *string       `json:"transaction-number"`
	PaymentMethodID     string        `json:"payment-method-id"`
	Included            []string      `json:"included"`
	PaymentMethod       PaymentMethod `json:"payment-method"`
	Links               Links         `json:"links"`
	Relationships       Relationships `json:"relationships"`
}

type ContributionData struct {
	ID         string                 `json:"id,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Attributes ContributionAttributes `json:"attributes"`
}

type Contribution struct {
	Data ContributionData `json:"data"`
}
