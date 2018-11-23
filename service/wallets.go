package service

import (
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
)

type Wallet struct {
	bc *blockchain.Blockchain
}

func NewWallet(bc *blockchain.Blockchain) *Wallet {
	return &Wallet{bc}
}

func (w *Wallet) ListAccounts(params string) (interface{}, error) {
	return w.bc.ListAccounts(params)
}

func (w *Wallet) GetAccount(params string) (interface{}, error) {
	return w.bc.GetAccount(params)
}

func (w *Wallet) GetBalanceByPrivateKey(privKey string) (interface{}, error) {
	return w.bc.GetBalanceByPrivateKey(privKey)
}

func (w *Wallet) GetListCustomTokenBalance(paymentAddress string) (*blockchain.ListCustomTokenBalance, error) {
	return w.bc.GetListCustomTokenBalance(paymentAddress)
}

func (w *Wallet) GetCoinAndCustomTokenBalance(privKey string, paymentAddress string) (*serializers.WalletBalances, error) {
	result := &serializers.WalletBalances{
		ListBalances: []serializers.WalletBalance{},
	}
	coinBalance, err := w.bc.GetBalanceByPrivateKey(privKey)
	if err != nil {
		return nil, err
	}
	listCustomTokenBalances, err := w.GetListCustomTokenBalance(paymentAddress)
	if err != nil {
		return nil, err
	}
	result.PaymentAddress = listCustomTokenBalances.Address
	// TODO check with order table of exchange
	inOrder := uint64(0)
	// end TODO
	balanceCoin := serializers.WalletBalance{
		TotalBalance:     coinBalance,
		SymbolCode:       "CONST",
		SymbolName:       "Constant",
		AvailableBalance: coinBalance - inOrder,
		ConstantValue:    0,
		InOrder:          inOrder,
	}
	result.ListBalances = append(result.ListBalances, balanceCoin)
	if len(listCustomTokenBalances.ListCustomTokenBalance) > 0 {
		for _, item := range listCustomTokenBalances.ListCustomTokenBalance {
			balanceCoin = serializers.WalletBalance{
				TotalBalance:     item.Amount,
				SymbolCode:       item.Symbol,
				SymbolName:       item.Name,
				AvailableBalance: item.Amount,
				ConstantValue:    0,
				InOrder:          0,
			}
		}
	}
	return result, nil
}

func (w *Wallet) Send(privKey string, req serializers.WalletSend) error {
	var err error
	switch req.Type {
	case 0:
		// send coin constant
		_, err = w.bc.Createandsendtransaction(privKey, req)
	case 1:
		// send coin constant
		err = w.bc.Sendcustomtokentransaction(privKey, req)
	}
	return err
}
