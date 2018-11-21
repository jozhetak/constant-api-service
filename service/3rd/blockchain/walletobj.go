package blockchain

type ListCustomTokenBalance struct {
	Address                string               `json:"PaymentAddress"`
	ListCustomTokenBalance []CustomTokenBalance `json:"ListCustomTokenBalance"`
}

type CustomTokenBalance struct {
	Name    string `json:"Name"`
	Symbol  string `json:"Symbol"`
	Amount  uint64 `json:"Amount"`
	TokenID string `json:"TokenID"`
}
