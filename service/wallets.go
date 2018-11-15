package service

import "github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"

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

func (w *Wallet) EncryptData(params string) (interface{}, error) {
	return w.bc.EncryptData(params)
}

func (w *Wallet) GetBalanceByPrivateKey(privKey string) (interface{}, error) {
	return w.bc.GetBalanceByPrivateKey(privKey)
}
