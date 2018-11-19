package serializers

type MarketResp struct {
	BaseCurrency         string `json:"BaseCurrency"`
	QuoteCurrency        string `json:"QuoteCurrency"`
	DisplayName          string `json:"DisplayName"`
	State                string `json:"State"`
	SymbolCode           string `json:"SymbolCode"`
	Icon                 string `json:"Icon"`
	TradeEnabled         bool   `json:"TradeEnabled"`
	FeePrecision         int    `json:"FeePrecision"`
	TradePricePrecision  int    `json:"TradePricePrecision"`
	TradeTotalPrecision  int    `json:"TradeTotalPrecision"`
	TradeAmountPrecision int    `json:"TradeAmountPrecision"`
}

type OrderResp struct {
	ID         uint    `json:"ID"`
	SymbolCode string  `json:"SymbolCode"`
	Price      float64 `json:"Price"`
	Quantity   uint    `json:"Quantity"`
	Type       string  `json:"Type"`
	Status     string  `json:"Status"`
	Side       string  `json:"Side"`
	Time       string  `json:"Time"`
}
