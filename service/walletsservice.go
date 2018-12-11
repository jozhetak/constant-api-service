package service

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/ninjadotorg/constant-api-service/dao/exchange"
	"github.com/ninjadotorg/constant-api-service/models"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
)

type WalletService struct {
	bc          *blockchain.Blockchain
	exchangeDAO *exchange.Exchange
}

func NewWalletService(bc *blockchain.Blockchain, ex *exchange.Exchange) *WalletService {
	return &WalletService{
		bc:          bc,
		exchangeDAO: ex,
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

func (w *WalletService) GetBalanceByPaymentAddress(paymentAddress string) (interface{}, error) {
	return w.bc.GetBalanceByPaymentAddress(paymentAddress)
}

func (w *WalletService) GetListCustomTokenBalance(paymentAddress string) (*blockchain.ListCustomTokenBalance, error) {
	return w.bc.GetListCustomTokenBalance(paymentAddress)
}

func (w *WalletService) GetCoinAndCustomTokenBalanceForUser(u *models.User) (*serializers.WalletBalances, error) {
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
	result.PaymentAddress = listCustomTokenBalances.PaymentAddress
	// get in order for constant
	inOrderConstant := uint64(0)
	oStatus := models.New
	oSide := models.Buy
	orders, err := w.exchangeDAO.OrderHistory("constantbond", &oStatus, &oSide, nil, nil, u)
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		inOrderConstant += order.Price * order.Quantity
	}
	balanceCoin := serializers.WalletBalance{
		TotalBalance:     coinBalance,
		SymbolCode:       strings.ToLower("CONST"),
		SymbolName:       "Constant",
		AvailableBalance: coinBalance - inOrderConstant,
		ConstantValue:    0,
		InOrder:          inOrderConstant,
		TokenID:          "",
	}
	result.ListBalances = append(result.ListBalances, balanceCoin)

	if len(listCustomTokenBalances.ListCustomTokenBalance) > 0 {
		for _, item := range listCustomTokenBalances.ListCustomTokenBalance {
			currency, err := w.exchangeDAO.FindCurrencyByToken(item.TokenID)
			if err != nil {
				return nil, errors.Wrapf(err, "w.exchangeService.FindCurrencyByToken %s", item.TokenID)
			}
			markets, err := w.exchangeDAO.ListMarkets(nil, currency)
			if err != nil {
				return nil, errors.Wrap(err, "w.exchangeService.ListMarkets")
			}

			oStatus := models.New
			oSide := models.Sell
			orders, _ := w.exchangeDAO.FindOrdersInMarkets(markets, &oStatus, &oSide)
			inOrderToken := uint64(0)
			for _, order := range orders {
				inOrderToken += order.Quantity
			}

			balanceCoin := serializers.WalletBalance{
				TotalBalance:     item.Amount,
				SymbolCode:       strings.ToLower(item.Symbol),
				SymbolName:       item.Name,
				AvailableBalance: item.Amount - inOrderToken,
				ConstantValue:    0,
				InOrder:          inOrderToken,
				TokenID:          item.TokenID,
			}
			result.ListBalances = append(result.ListBalances, balanceCoin)
		}
	}
	return result, nil
}

func (w *WalletService) GetCoinAndCustomTokenBalanceForPaymentAddress(paymentAddress string) (*serializers.WalletBalances, error) {
	result := &serializers.WalletBalances{
		ListBalances: []serializers.WalletBalance{},
	}
	coinBalance, err := w.bc.GetBalanceByPaymentAddress(paymentAddress)
	if err != nil {
		return nil, err
	}
	listCustomTokenBalances, err := w.GetListCustomTokenBalance(paymentAddress)
	if err != nil {
		return nil, err
	}
	result.PaymentAddress = listCustomTokenBalances.PaymentAddress
	// end
	// get in order for
	balanceCoin := serializers.WalletBalance{
		TotalBalance:     coinBalance,
		SymbolCode:       strings.ToLower("CONST"),
		SymbolName:       "Constant",
		AvailableBalance: coinBalance,
		ConstantValue:    0,
		InOrder:          0,
	}
	result.ListBalances = append(result.ListBalances, balanceCoin)
	if len(listCustomTokenBalances.ListCustomTokenBalance) > 0 {
		for _, item := range listCustomTokenBalances.ListCustomTokenBalance {
			balanceCoin = serializers.WalletBalance{
				TotalBalance:     item.Amount,
				SymbolCode:       strings.ToLower(item.Symbol),
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
