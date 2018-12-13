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
	PartyID        uint        `json:"PartyID"`
	PaymentAddress string      `json:"PaymentAddress"`
	PaymentForm    PaymentForm `json:"PaymentForm"`
	Amount         float64     `json:"Amount"`
}

type PaymentForm struct {
	RoutingNumber string `json:"RoutingNumber"`

	ContactEmail string `json:"ContactEmail"` // required
	ContactName  string `json:"ContactName"`  // required

	AchCheckType      AchCheckType `json:"AchCheckType"`      // (enum): personal or business.
	BankAccountName   string       `json:"BankAccountName"`   // required in ACH
	BankAccountType   string       `json:"BankAccountType"`   // required in ACH
	BankAccountNumber string       `json:"BankAccountNumber"` // required in ACH
	BankName          string       `json:"BankName"`          // required in ACH
}

type PrimetrustChangedRequest struct {
	ID           string                 `json:"id"`
	AccountID    string                 `json:"account_id"`
	Action       string                 `json:"action"`
	Data         map[string]interface{} `json:"data"`
	ResourceID   string                 `json:"resource_id"`
	ResourceType string                 `json:"resource_type"`
}
