package serializers

type WalletBalances struct {
	Address      string
	ListBalances []WalletBalance
}

type WalletBalance struct {
	SymbolName       string
	SymbolCode       string
	TotalBalance     uint64
	AvailableBalance uint64
	InOrder          uint64
	ConstantValue    uint64
}
