package models

const (
	DisbursementType = "disbursements"
)

type DisbursementAttributes struct {
	ID                string                  `json:"id,omitempty"`
	AccountID         string                  `json:"account-id"`
	Amount            float64                 `json:"amount"`
	CustomerReference string                  `json:"customer-reference"`
	Description       string                  `json:"description"`
	PaymentMethodID   string                  `json:"payment-method-id"`
	PaymentMethod     PaymentMethodAttributes `json:"payment-method"`
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

type DisbursementsResponse struct {
	CollectionResponse
	Data []DisbursementData `json:"data"`
}
