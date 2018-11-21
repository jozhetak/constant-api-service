package serializers

type BorrowReq struct {
	Amount         int64   `json:"Amount" binding:"required"`
	HashKey        string  `json:"HashKey" binding:"required"`
	CollateralTxID string  `json:"CollateralTxID" binding:"required"`
	Collateral     string  `json:"Collateral"`
	StartDate      string  `json:"StartDate"`
	EndDate        string  `json:"EndDate"`
	Rate           float64 `json:"Rate"`
}
