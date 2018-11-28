package models

import (
	"time"
	"github.com/mongodb/mongo-go-driver/bson/decimal"
)

const (
	ContributionType             = "contributions"
	ContributionPaymentTypeCheck = "check"
	ContributionPaymentTypeWire  = "wire"
)

type ContributionAttributes struct {
	AccountID           string                  `json:"account-id"`
	Amount              decimal.Decimal128      `json:"amount"`
	ContactEmail        string                  `json:"contact-email"`
	ContactName         string                  `json:"contact-name"`
	ContactID           string                  `json:"contact-id"`
	Message             *string                 `json:"message"`
	PaymentMethod       PaymentMethodAttributes `json:"payment-method"`
	PaymentMethodID     string                  `json:"payment-method-id"`
	PaymentType         string                  `json:"payment-type"`
	SavePaymentMethod   bool                    `json:"save-payment-method"`
	SpecialInstructions *string                 `json:"special-instructions"`
	SurplusToParent     bool                    `json:"surplus_to_parent"`
	ShortageToParent    bool                    `json:"shortage_to_parent"`
	ID                  string                  `json:"id"`
	AmountExpected      decimal.Decimal128      `json:"amount-expected"`
	ParentID            string                  `json:"parent_id"`
	PrimaryChild        bool                    `json:"primary_child"`
	PaymentDetails      *string                 `json:"payment-details"`
	ReferenceNumber     string                  `json:"reference-number"`
	SpecialType         *string                 `json:"special-type"`
	CreatedAt           *time.Time              `json:"created-at"`
	Status              string                  `json:"status"`
	TransactionNumber   *string                 `json:"transaction-number"`
}

type ContributionData struct {
	ID            string                 `json:"id,omitempty"`
	Type          string                 `json:"type,omitempty"`
	Attributes    ContributionAttributes `json:"attributes"`
	Links         Links                  `json:"links"`
	Relationships Relationships          `json:"relationships"`
	Included      []string               `json:"included"`
}

type Contribution struct {
	Data ContributionData `json:"data"`
}

type ContributionResponse struct {
	CollectionResponse
	Data []ContributionData `json:"data"`
}
