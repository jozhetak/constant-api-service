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
	PaymentType   PaymentType `json:"PaymentType"` // required (enum, sql:true): ach, check, credit_card, wire, or wire_international
	RoutingNumber string      `json:"RoutingNumber"`

	ContactEmail string `json:"ContactEmail"` // required
	ContactName  string `json:"ContactName"`  // required
	ContactID    string `json:"ContactID"`

	AchCheckType    AchCheckType `json:"AchCheckType"`    // (enum): personal or business.
	BankAccountName string       `json:"BankAccountName"` // required in ACH
	BankAccountType string       `json:"BankAccountType"` // required in ACH
	BankName        string       `json:"BankName"`        // required in ACH

	CreditCardCvv            string `json:"CreditCardCvv"`            // required in CC
	CreditCardExpirationDate string `json:"CreditCardExpirationDate"` // required in CC
	CreditCardNumber         string `json:"CreditCardNumber"`         // required in CC
	CreditCardName           string `json:"CreditCardName"`
	CreditCardPostalCode     string `json:"CreditCardPostalCode"`
	CreditCardType           string `json:"CreditCardType"`

	Last4 string `json:"Last4"`
}

type PrimetrustChangedRequest struct {
	ID           string                 `json:"id"`
	AccountID    string                 `json:"account_id"`
	Action       string                 `json:"action"`
	Data         map[string]interface{} `json:"data"`
	ResourceID   string                 `json:"resource_id"`
	ResourceType string                 `json:"resource_type"`
}
