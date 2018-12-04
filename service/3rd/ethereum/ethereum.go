package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ninjadotorg/constant-api-service/conf"
)

type EthereumService struct {
	EthChainEnpoint string
	SimpleLoanOwner string
	SimpleLoadAddr  string
	client          *ethclient.Client
}

func Init(conf *config.Config) *EthereumService {
	service := &EthereumService{
		EthChainEnpoint: conf.EthChainEndpoint,
		SimpleLoanOwner: conf.SimpleLoanOwner,
		SimpleLoadAddr:  conf.SimpleLoanAddr,
	}
	return service
}

func (s *EthereumService) GetClient() (*ethclient.Client, error) {
	if s.client != nil {
		return s.client, nil
	}

	client, err := ethclient.Dial(s.EthChainEnpoint)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *EthereumService) GetGasPrice() (*big.Int, error) {
	client, err := s.GetClient()
	if err != nil {
		return big.NewInt(0), err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return big.NewInt(0), err
	}

	return gasPrice, nil
}

func (s *EthereumService) SendSignedTransaction(fromPrvKey string, to string, value *big.Int, data []byte) (string, error) {
	client, err := s.GetClient()

	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(fromPrvKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasLimit := uint64(300000) // in units
	gasPrice, err := s.GetGasPrice()
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(to)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	ts := &(types.Transactions{signedTx})
	rawTxBytes := ts.GetRlp(0)

	rtx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &rtx)
	err = client.SendTransaction(context.Background(), rtx)
	if err != nil {
		return "", err
	}
	return rtx.Hash().Hex(), nil
}

func (s *EthereumService) GetSimpleLoanAbi() abi.ABI {
	simpleLoanAbi, _ := abi.JSON(strings.NewReader(SimpleLoanABI))
	return simpleLoanAbi
}

func (s *EthereumService) SimpleLoanAcceptLoan(lid string, key string, offchain string) (string, error) {
	simpleLoanAbi := s.GetSimpleLoanAbi()

	bytesData, _ := simpleLoanAbi.Pack("AcceptLoan", lid, key, offchain)

	txHash, err := s.SendSignedTransaction(s.SimpleLoanOwner, s.SimpleLoadAddr, big.NewInt(0), bytesData)

	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (s *EthereumService) SimpleLoanRejectLoan(lid string, offchain string) (string, error) {
	simpleLoanAbi := s.GetSimpleLoanAbi()

	bytesData, _ := simpleLoanAbi.Pack("RejectLoan", lid, offchain)

	txHash, err := s.SendSignedTransaction(s.SimpleLoanOwner, s.SimpleLoadAddr, big.NewInt(0), bytesData)

	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (s *EthereumService) SimpleLoanAddPayment(lid string, amount *big.Int, offchain string) (string, error) {
	simpleLoanAbi := s.GetSimpleLoanAbi()

	bytesData, _ := simpleLoanAbi.Pack("RejectLoan", lid, offchain)

	txHash, err := s.SendSignedTransaction(s.SimpleLoanOwner, s.SimpleLoadAddr, big.NewInt(0), bytesData)

	if err != nil {
		return "", err
	}

	return txHash, nil
}

//func Demo() {
//	key, _ := crypto.GenerateKey()
//	auth := bind.NewKeyedTransactor(key)
//	alloc := make(core.GenesisAlloc)
//	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
//	sim := backends.NewSimulatedBackend(alloc, 1000)
//
//	hash := common.HexToHash("")
//	scAddress := common.BytesToAddress(hash[:])
//	simpleLoan, _ := NewSimpleLoan(scAddress, sim)
//	opts := &bind.TransactOpts{
//
//	}
//	simpleLoan.AcceptLoan(opts, [32]byte{}, [32]byte{}, [32]byte{})
//}
