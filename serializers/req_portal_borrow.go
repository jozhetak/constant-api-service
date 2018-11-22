package serializers

type BorrowReq struct {
	Amount         int64       `json:"Amount" binding:"required"`
	HashKey        string      `json:"HashKey" binding:"required"`
	CollateralTxID string      `json:"CollateralTxID" binding:"required"`
	Collateral     string      `json:"Collateral"`
	StartDate      string      `json:"StartDate"`
	EndDate        string      `json:"EndDate"`
	Rate           float64     `json:"Rate"`
	PaymentAddress string      `json:"PaymentAddress"`
	LoanRequest    LoanRequest `json:"LoanRequest"`
}

type LoanParams struct {
	InterestRate     uint64 `json:"InterestRate"` // basis points, e.g. 125 represents 1.25%
	Maturity         uint64 `json:"Maturity"`     // seconds
	LiquidationStart uint64 `json:"Maturity"`     // ratio between collateral and debt to start auto-liquidation, stored in basis points
}

type LoanRequest struct {
	Params           LoanParams `json:"Params"`
	LoanID           string     `json:"LoanID"` // 32 bytes
	CollateralType   string     `json:"CollateralType"`
	CollateralTx     string     `json:"CollateralTx"` // Tx hash in case of ETH
	CollateralAmount string     `json:"CollateralAmount"`

	LoanAmount     uint64 `json:"LoanAmount"`
	ReceiveAddress string `json:"ReceiveAddress"`

	KeyDigest string `json:"KeyDigest"` // 32 bytes, from sha256
}
