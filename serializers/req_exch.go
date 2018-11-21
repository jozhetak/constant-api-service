package serializers

type OrderReq struct {
	Symbol   string `json:"Symbol" binding:"required"`
	Price    uint64 `json:"Price" binding:"required"`
	Quantity uint64 `json:"Quantity" binding:"required"`
	Type     string `json:"Type" binding:"required"`
	Side     string `json:"Side" binding:"required"`
}

type OrderMsg struct {
	ID     int    `json:"id"`
	Price  uint64 `json:"price"`
	Size   uint64 `json:"size"`
	Side   string `json:"side"`
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
}

type OrderPubMsg struct {
	Type  string    `json:"type"`
	Order *OrderMsg `json:"order"`
}
