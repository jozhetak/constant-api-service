package serializers

type WalletBalances struct {
	PaymentAddress string
	ListBalances   []WalletBalance
}

type WalletBalance struct {
	SymbolName       string
	SymbolCode       string
	Token            string
	TotalBalance     uint64
	AvailableBalance uint64
	InOrder          uint64
	ConstantValue    uint64
}
