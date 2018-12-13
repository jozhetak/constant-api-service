package models

const (
	FundsTransferMethodType                               = "payment-methods"
	FundsTransferMethodAchCheckTypePersonal               = "personal"
	FundsTransferMethodAchCheckTypeBusiness               = "business"
	FundsTransferMethodBankAccountTypeChecking            = "checking"
	FundsTransferMethodBankAccountTypeSavings             = "savings"
	FundsTransferMethodFundsTransferTypeAch               = "ach"
	FundsTransferMethodFundsTransferTypeCheck             = "check"
	FundsTransferMethodFundsTransferTypeCreditCard        = "credit_card"
	FundsTransferMethodFundsTransferTypeWire              = "wire"
	FundsTransferMethodFundsTransferTypeWireInternational = "wire_international"
	FundsTransferMethodCreditCardTypeMC                   = "MC"
	FundsTransferMethodCreditCardTypeVI                   = "VI"
)

type FundsTransferMethodAttributes struct {
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
	FundsTransferType         string `json:"funds-transfer-type"` // (enum, sql:true): ach, check, credit_card, wire, or wire_international
	RoutingNumber             string `json:"routing-number,omitempty"`
	SwiftCode                 string `json:"swift-code,omitempty"`
	IpAddress                 string `json:"ip-address,omitempty"`
	BankAccountNumber         string `json:"bank-account-number,omitempty"`
	CreditCardCvv             string `json:"credit-card-cvv,omitempty"`
	CreditCardExpirationDate  string `json:"credit-card-expiration-date,omitempty"`
	CreditCardNumber          string `json:"credit-card-number,omitempty"`
}

type FundsTransferMethodData struct {
	ID            string                        `json:"id,omitempty"`
	Type          string                        `json:"type,omitempty"`
	Attributes    FundsTransferMethodAttributes `json:"attributes"`
	Links         Links                         `json:"links,omitempty"`
	Relationships Relationships                 `json:"relationships,omitempty"`
}

type FundsTransferMethod struct {
	Data FundsTransferMethodData `json:"data"`
}

type FundsTransferMethodsResponse struct {
	CollectionResponse
	Data []FundsTransferMethodData `json:"data"`
}
