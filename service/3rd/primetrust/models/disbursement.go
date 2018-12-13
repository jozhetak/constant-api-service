package models

const (
	DisbursementType = "disbursements"
)

type DisbursementAttributes struct {
	ID                    string                         `json:"id,omitempty"`
	AccountID             string                         `json:"account-id,omitempty"`
	Amount                float64                        `json:"amount,omitempty"`
	CustomerReference     string                         `json:"customer-reference,omitempty"`
	Description           string                         `json:"description,omitempty"`
	FundsTransferMethodID string                         `json:"funds-transfer-method-id,omitempty"`
	FundsTransferMethod   *FundsTransferMethodAttributes `json:"funds-transfer-method,omitempty"`
	SpecialType           string                         `json:"special-type,omitempty"`
	Status                string                         `json:"status,omitempty"`
	ContactEmail          string                         `json:"contact-email,omitempty"`
	ContactName           string                         `json:"contact-name,omitempty"`
}

type DisbursementData struct {
	ID            string                  `json:"id,omitempty"`
	Type          string                  `json:"type,omitempty"`
	Attributes    *DisbursementAttributes `json:"attributes,omitempty"`
	Links         *Links                  `json:"links,omitempty"`
	Relationships *Relationships          `json:"relationships,omitempty"`
}

type Disbursement struct {
	Data *DisbursementData `json:"data"`
}

type DisbursementsResponse struct {
	CollectionResponse
	Data []DisbursementData `json:"data"`
}
