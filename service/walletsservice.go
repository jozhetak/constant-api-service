package service

import (
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
)

type WalletService struct {
	bc              *blockchain.Blockchain
	exchangeService *ExchangeService
}

func NewWalletService(bc *blockchain.Blockchain, ex *ExchangeService) *WalletService {
	return &WalletService{
		bc:              bc,
		exchangeService: ex,
	}
}

func (w *WalletService) ListAccounts(params string) (interface{}, error) {
	return w.bc.ListAccounts(params)
}

func (w *WalletService) GetAccount(params string) (interface{}, error) {
	return w.bc.GetAccount(params)
}

func (w *WalletService) GetBalanceByPrivateKey(privKey string) (interface{}, error) {
	return w.bc.GetBalanceByPrivateKey(privKey)
}

func (w *WalletService) GetListCustomTokenBalance(paymentAddress string) (*blockchain.ListCustomTokenBalance, error) {
	return w.bc.GetListCustomTokenBalance(paymentAddress)
}

func (w *WalletService) GetCoinAndCustomTokenBalance(u *models.User) (*serializers.WalletBalances, error) {
	result := &serializers.WalletBalances{
		ListBalances: []serializers.WalletBalance{},
	}
	coinBalance, err := w.bc.GetBalanceByPrivateKey(u.PrivKey)
	if err != nil {
		return nil, err
	}
	listCustomTokenBalances, err := w.GetListCustomTokenBalance(u.PaymentAddress)
	if err != nil {
		return nil, err
	}
	result.PaymentAddress = listCustomTokenBalances.Address
	// get in order for constant
	inOrderConstant := uint64(0)
	orders, _ := w.exchangeService.UserOrderHistory(u, "constantbond", "new", "", nil, nil)
	for _, order := range orders {
		inOrderConstant += order.Price * order.Quantity
	}
	// end
	// get in order for
	balanceCoin := serializers.WalletBalance{
		TotalBalance:     coinBalance,
		SymbolCode:       "CONST",
		SymbolName:       "Constant",
		AvailableBalance: coinBalance - inOrderConstant,
		ConstantValue:    0,
		InOrder:          inOrderConstant,
	}
	result.ListBalances = append(result.ListBalances, balanceCoin)
	if len(listCustomTokenBalances.ListCustomTokenBalance) > 0 {
		for _, item := range listCustomTokenBalances.ListCustomTokenBalance {
			// TODO
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

func (w *WalletService) Send(privKey string, req serializers.WalletSend) error {
	var err error
	switch req.Type {
	case 0:
		// send coin constant
		_, err = w.bc.CreateAndSendConstantTransaction(privKey, req)
	case 1:
		// send coin constant
		err = w.bc.SendCustomTokenTransaction(privKey, req)
	}
	return err
}
