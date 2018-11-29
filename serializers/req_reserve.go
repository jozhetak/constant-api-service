package serializers

type ReserveContributionRequest struct {
	PartyID        int         `json:"PartyID"`
	PaymentAddress string      `json:"PaymentAddress"`
	PaymentForm    PaymentForm `json:"PaymentForm"`
}

type ReserveDisbursementRequest struct {
	PartyID     int         `json:"PartyID"`
	PaymentForm PaymentForm `json:"PaymentForm"`
}

type PaymentForm struct {
	PaymentType   string `json:"PaymentType"` // required (enum, sql:true): ach, check, credit_card, wire, or wire_international
	RoutingNumber string `json:"RoutingNumber"`

	ContactEmail string `json:"ContactEmail"` // required
	ContactName  string `json:"ContactName"`  // required
	ContactID    string `json:"ContactID"`

	AchCheckType    string `json:"AchCheckType"`    // (enum): personal or business.
	BankAccountName string `json:"BankAccountName"` // required in ACH
	BankAccountType string `json:"BankAccountType"` // required in ACH
	BankName        string `json:"BankName"`        // required in ACH

	CreditCardCvv            string `json:"credit-card-cvv"`             // required in CC
	CreditCardExpirationDate string `json:"credit-card-expiration-date"` // required in CC
	CreditCardNumber         string `json:"credit-card-number"`          // required in CC
	CreditCardName           string `json:"CreditCardName"`
	CreditCardPostalCode     string `json:"CreditCardPostalCode"`
	CreditCardType           string `json:"CreditCardType"`

	Last4 string `json:"last-4"`
}
