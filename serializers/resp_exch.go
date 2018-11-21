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
	ID         uint   `json:"ID"`
	SymbolCode string `json:"SymbolCode"`
	Price      uint64 `json:"Price"`
	Quantity   uint64 `json:"Quantity"`
	Type       string `json:"Type"`
	Status     string `json:"Status"`
	Side       string `json:"Side"`
	Time       string `json:"Time"`
}

type SymbolRate struct {
	SymbolCode string `json:"SymbolCode"`
	Volume     uint64 `json:"Volume"`
	Last       uint64 `json:"Last"`
	High       uint64 `json:"High"`
	Low        uint64 `json:"Low"`
	PrevPrice  uint64 `json:"PrevPrice"`
	PrevVolume uint64 `json:"PrevVolume"`
}

type MarketRate struct {
	SymbolCode string `json:"SymbolCode"`
	Last       uint64 `json:"Last"`
	Bid        uint64 `json:"Bid"`
	Ask        uint64 `json:"Ask"`
	Volume     uint64 `json:"Volume"`
	High24h    uint64 `json:"24hHigh"`
	Low24h     uint64 `json:"24hLow"`
}
