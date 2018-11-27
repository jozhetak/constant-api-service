package serializers

type BorrowReq struct {
	StartDate   string      `json:"StartDate"`
	LoanRequest LoanRequest `json:"LoanRequest"`
}

type LoanParams struct {
	InterestRate     uint64 `json:"InterestRate"`     // basis points, e.g. 125 represents 1.25%
	Maturity         uint64 `json:"Maturity"`         // seconds
	LiquidationStart uint64 `json:"LiquidationStart"` // ratio between collateral and debt to start auto-liquidation, stored in basis points
}

type LoanRequest struct {
	Params           LoanParams `json:"Params"`
	LoanID           string     `json:"LoanID"` // 32 bytes
	CollateralType   string     `json:"CollateralType"`
	CollateralAmount string     `json:"CollateralAmount"`

	LoanAmount     uint64 `json:"LoanAmount"`
	ReceiveAddress string `json:"ReceiveAddress"`

	KeyDigest string `json:"KeyDigest"` // 32 bytes, from sha256
}

type LoanWithdraw struct {
	LoanID string
	Key    string
}

type LoanPayment struct {
	LoanID string
}
