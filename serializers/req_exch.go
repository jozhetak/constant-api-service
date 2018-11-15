package serializers

type OrderReq struct {
	Symbol   string  `json:"Symbol" binding:"required"`
	Price    float64 `json:"Price" binding:"required"`
	Quantity uint    `json:"Quantity" binding:"required"`
	Type     string  `json:"Type" binding:"required"`
	Side     string  `json:"Side" binding:"required"`
}
