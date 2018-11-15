package serializers

type BorrowReq struct {
	Amount         int64  `json:"Amount" binding:"required"`
	Hash           string `json:"Hash" binding:"required"`
	TxID           string `json:"TxID" binding:"required"`
	PaymentAddress string `json:"PaymentAddress" binding:"required"`
}
