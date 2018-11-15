package serializers

type BorrowResp struct {
	ID             uint   `json:"ID"`
	Amount         int64  `json:"Amount"`
	Hash           string `json:"Hash"`
	TxID           string `json:"TxID"`
	PaymentAddress string `json:"PaymentAddress"`
	State          string `json:"State"`
	CreatedAt      string `json:"CreatedAt"`
}
