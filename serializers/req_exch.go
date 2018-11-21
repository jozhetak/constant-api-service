package serializers

type OrderReq struct {
	Symbol   string `json:"Symbol" binding:"required"`
	Price    uint64 `json:"Price" binding:"required"`
	Quantity uint64 `json:"Quantity" binding:"required"`
	Type     string `json:"Type" binding:"required"`
	Side     string `json:"Side" binding:"required"`
}
