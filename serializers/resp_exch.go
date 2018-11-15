package serializers

type MarketResp struct {
	BaseCurrency   string `json:"BaseCurrency"`
	MarketCurrency string `json:"MarketCurrency"`
	Symbol         string `json:"Symbol"`
}

type OrderResp struct {
	ID       uint    `json:"ID"`
	Symbol   string  `json:"Symbol"`
	Price    float64 `json:"Price"`
	Quantity uint    `json:"Quantity"`
	Type     string  `json:"Type"`
	Status   string  `json:"Status"`
	Side     string  `json:"Side"`
	Time     string  `json:"Time"`
}
