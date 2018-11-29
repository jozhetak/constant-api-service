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
	ContactID                 string `json:"contact-id"`
	AchCheckType              string `json:"ach-check-type"` // (enum): personal or business.
	BankAccountName           string `json:"bank-account-name"`
	BankAccountType           string `json:"bank-account-type"`
	BankName                  string `json:"bank-name"`
	CheckPayee                string `json:"check-payee"`
	ContactEmail              string `json:"contact-email"`
	ContactName               string `json:"contact-name"`
	CreditCardName            string `json:"credit-card-name"`
	CreditCardPostalCode      string `json:"credit-card-postal-code"`
	CreditCardType            string `json:"credit-card-type"`
	Inactive                  bool   `json:"inactive"`
	IntermediaryBankName      string `json:"intermediary-bank-name"`
	IntermediaryBankReference string `json:"intermediary-bank-reference"`
	Last4                     string `json:"last-4"`
	PaymentType               string `json:"payment-type"` // (enum, sql:true): ach, check, credit_card, wire, or wire_international
	RoutingNumber             string `json:"routing-number"`
	SwiftCode                 string `json:"swift-code"`
	IpAddress                 string `json:"ip-address"`
	BankAccountNumber         string `json:"bank-account-number"`
	CreditCardCvv             string `json:"credit-card-cvv"`
	CreditCardExpirationDate  string `json:"credit-card-expiration-date"`
	CreditCardNumber          string `json:"credit-card-number"`
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
