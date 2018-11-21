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

type TransactionDetail struct {
	BlockHash string `json:"BlockHash"`
	Index     uint64 `json:"Index"`
	ChainId   byte   `json:"ChainId"`
	Hash      string `json:"Hash"`
	Version   int8   `json:"Version"`
	Type      string `json:"Type"` // Transaction type
	LockTime  int64  `json:"LockTime"`
	Fee       uint64 `json:"Fee"` // Fee applies: always consant

	Descs    []interface{} `json:"Descs"`
	JSPubKey []byte        `json:"JSPubKey,omitempty"` // 64 bytes
	JSSig    []byte        `json:"JSSig,omitempty"`    // 64 bytes

	AddressLastByte byte   `json:"AddressLastByte"`
	MetaData        string `json:"MetaData"`
}
