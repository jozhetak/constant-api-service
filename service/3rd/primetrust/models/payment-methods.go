package models

const (
	PaymentMethodType                         = "payment-methods"
	PaymentMethodAchCheckTypePersonal         = "personal"
	PaymentMethodAchCheckTypeBusiness         = "business"
	PaymentMethodBankAccountTypeChecking      = "checking"
	PaymentMethodBankAccountTypeSavings       = "savings"
	PaymentMethodPaymentTypeAch               = "ach"
	PaymentMethodPaymentTypeCheck             = "check"
	PaymentMethodPaymentTypeCreditCard        = "credit_card"
	PaymentMethodPaymentTypeWire              = "wire"
	PaymentMethodPaymentTypeWireInternational = "wire_international"
	PaymentMethodCreditCardTypeMC             = "MC"
	PaymentMethodCreditCardTypeVI             = "VI"
)

type PaymentMethodAttributes struct {
	ID                        string `json:"id,omitempty"`
	ContactID                 string `json:"contact-id,omitempty"`
	AchCheckType              string `json:"ach-check-type,omitempty"` // (enum): personal or business.
	BankAccountName           string `json:"bank-account-name,omitempty"`
	BankAccountType           string `json:"bank-account-type,omitempty"`
	BankName                  string `json:"bank-name,omitempty"`
	CheckPayee                string `json:"check-payee,omitempty"`
	ContactEmail              string `json:"contact-email,omitempty"`
	ContactName               string `json:"contact-name,omitempty"`
	CreditCardName            string `json:"credit-card-name,omitempty"`
	CreditCardPostalCode      string `json:"credit-card-postal-code,omitempty"`
	CreditCardType            string `json:"credit-card-type,omitempty"`
	Inactive                  bool   `json:"inactive,omitempty"`
	IntermediaryBankName      string `json:"intermediary-bank-name,omitempty"`
	IntermediaryBankReference string `json:"intermediary-bank-reference,omitempty"`
	Last4                     string `json:"last-4,omitempty"`
	PaymentType               string `json:"payment-type"` // (enum, sql:true): ach, check, credit_card, wire, or wire_international
	RoutingNumber             string `json:"routing-number,omitempty"`
	SwiftCode                 string `json:"swift-code,omitempty"`
	IpAddress                 string `json:"ip-address,omitempty"`
	BankAccountNumber         string `json:"bank-account-number,omitempty"`
	CreditCardCvv             string `json:"credit-card-cvv,omitempty"`
	CreditCardExpirationDate  string `json:"credit-card-expiration-date,omitempty"`
	CreditCardNumber          string `json:"credit-card-number,omitempty"`
}

type PaymentMethodData struct {
	ID            string                  `json:"id,omitempty"`
	Type          string                  `json:"type,omitempty"`
	Attributes    PaymentMethodAttributes `json:"attributes"`
	Links         Links                   `json:"links"`
	Relationships Relationships           `json:"relationships"`
}

type PaymentMethod struct {
	Data PaymentMethodData `json:"data"`
}

type PaymentMethodsResponse struct {
	CollectionResponse
	Data []PaymentMethodData `json:"data"`
}
