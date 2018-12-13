package models

import (
	"time"
)

const (
	ContributionType             = "contributions"
	ContributionPaymentTypeCheck = "check"
	ContributionPaymentTypeWire  = "wire"
)

type ContributionAttributes struct {
	ID                      string                         `json:"id,omitempty"`
	AccountID               string                         `json:"account-id,omitempty"`
	Amount                  float64                        `json:"amount,omitempty"`
	ContactEmail            string                         `json:"contact-email,omitempty"`
	ContactName             string                         `json:"contact-name,omitempty"`
	ContributorEmail        string                         `json:"contributor-email,omitempty"`
	ContributorName         string                         `json:"contributor-name,omitempty"`
	ContactID               string                         `json:"contact-id,omitempty"`
	Message                 string                         `json:"message,omitemtpy"`
	FundsTransferMethod     *FundsTransferMethodAttributes `json:"funds-transfer-method,omitempty"`
	FundsTransferMethodID   string                         `json:"funds-transfer-method-id,omitempty"`
	PaymentType             string                         `json:"payment-type,omitempty"`
	SaveFundsTransferMethod bool                           `json:"save-funds-transfer-method,omitempty"`
	SpecialInstructions     string                         `json:"special-instructions,omitempty"`
	SurplusToParent         bool                           `json:"surplus_to_parent,omitempty"`
	ShortageToParent        bool                           `json:"shortage_to_parent,omitempty"`
	AmountExpected          float64                        `json:"amount-expected,omitempty"`
	ParentID                string                         `json:"parent_id,omitempty"`
	PrimaryChild            bool                           `json:"primary_child,omitempty"`
	PaymentDetails          string                         `json:"payment-details,omitempty"`
	ReferenceNumber         string                         `json:"reference-number,omitempty"`
	SpecialType             string                         `json:"special-type,omitempty"`
	CreatedAt               time.Time                      `json:"created-at,omitempty"`
	Status                  string                         `json:"status,omitempty"`
	TransactionNumber       string                         `json:"transaction-number,omitempty"`
}

type ContributionData struct {
	ID            string                  `json:"id,omitempty"`
	Type          string                  `json:"type,omitempty"`
	Attributes    *ContributionAttributes `json:"attributes"`
	Links         *Links                  `json:"links,omitempty"`
	Relationships *Relationships          `json:"relationships,omitempty"`
	Included      []string                `json:"included,omitempty"`
}

type Contribution struct {
	Data *ContributionData `json:"data"`
}

type ContributionsResponse struct {
	CollectionResponse
	Data []ContributionData `json:"data"`
}
