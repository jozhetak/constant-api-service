package serializers

type BorrowResp struct {
	ID             uint    `json:"ID"`
	Amount         int64   `json:"Amount"`
	Hash           string  `json:"HashKey"`
	CollateralTxID string  `json:"CollateralTxID"`
	StartDate      string  `json:"StartDate"`
	EndDate        string  `json:"EndDate"`
	Collateral     string  `json:"Collateral"`
	Rate           float64 `json:"Rate"`
	State          string  `json:"State"`
	CreatedAt      string  `json:"CreatedAt"`
}
