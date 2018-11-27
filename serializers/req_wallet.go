package serializers

type WalletSend struct {
	Type             int // 0: constant, 1: token
	TokenID          string
	TokenName        string
	TokenSymbol      string
	PaymentAddresses map[string]uint64
}
