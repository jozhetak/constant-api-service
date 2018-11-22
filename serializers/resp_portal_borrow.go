package serializers

type BorrowResp struct {
	ID                  uint `json:"ID"`
	PaymentAddress      string
	LoanAmount          int64
	KeyDigest           string
	LoanID              string
	CollateralType      string
	CollateralAmount    string
	StartDate           string
	EndDate             string
	InterestRate        int64
	Maturity            int64
	LiquidationStart    int64
	ConstantLoanTxID    string
	ConstantPaymentTxID string
	State               string
	CreatedAt           string
}
