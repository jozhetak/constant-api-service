package serializers

type PaymentType string

const (
	PaymentTypeAch               PaymentType = "ach"
	PaymentTypeCheck             PaymentType = "check"
	PaymentTypeCreditCard        PaymentType = "credit_card"
	PaymentTypeWire              PaymentType = "wire"
	PaymentTypeWireInternational PaymentType = "wire_international"
)

type AchCheckType string

const (
	AchCheckTypePersonal AchCheckType = "personal"
	AchCheckTypeBusiness AchCheckType = "business"
)

type ReserveContributionRequest struct {
	PartyID        uint        `json:"PartyID"`
	PaymentAddress string      `json:"PaymentAddress"`
	PaymentForm    PaymentForm `json:"PaymentForm"`
	Amount         float64     `json:"Amount"`
}

type ReserveDisbursementRequest struct {
	PartyID     uint        `json:"PartyID"`
	PaymentForm PaymentForm `json:"PaymentForm"`
	Amount      float64     `json:"Amount"`
}

type PaymentForm struct {
	PaymentType   PaymentType `json:"PaymentType,omitempty"` // required (enum, sql:true): ach, check, credit_card, wire, or wire_international
	RoutingNumber string      `json:"RoutingNumber,omitempty"`

	ContactEmail string `json:"ContactEmail,omitempty"` // required
	ContactName  string `json:"ContactName,omitempty"`  // required
	ContactID    string `json:"ContactID,omitempty"`

	AchCheckType      AchCheckType `json:"AchCheckType,omitempty"`      // (enum): personal or business.
	BankAccountName   string       `json:"BankAccountName,omitempty"`   // required in ACH
	BankAccountType   string       `json:"BankAccountType,omitempty"`   // required in ACH
	BankAccountNumber string       `json:"BankAccountNumber,omitempty"` // required in ACH
	BankName          string       `json:"BankName,omitempty"`          // required in ACH

	CreditCardCvv            string `json:"CreditCardCvv,omitempty"`            // required in CC
	CreditCardExpirationDate string `json:"CreditCardExpirationDate,omitempty"` // required in CC
	CreditCardNumber         string `json:"CreditCardNumber,omitempty"`         // required in CC
	CreditCardName           string `json:"CreditCardName,omitempty"`
	CreditCardPostalCode     string `json:"CreditCardPostalCode,omitempty"`
	CreditCardType           string `json:"CreditCardType,omitempty"`

	Last4 string `json:"Last4,omitempty"`
}

type PrimetrustChangedRequest struct {
	ID           string                 `json:"id"`
	AccountID    string                 `json:"account_id"`
	Action       string                 `json:"action"`
	Data         map[string]interface{} `json:"data"`
	ResourceID   string                 `json:"resource_id"`
	ResourceType string                 `json:"resource_type"`
}
