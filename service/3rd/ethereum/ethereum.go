package ethereum

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ninjadotorg/constant-api-service/conf"
	"github.com/sendgrid/sendgrid-go"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core"
)

type EthereumClient struct {
	SimpleLoanOwner		string
	SimpleLoadAddr		string
}

func Init(conf *config.Config) *EthereumClient {
	client := &EthereumClient{conf.SimpleLoanOwner, conf.SimpleLoanAddress}
	return client
}

func Demo() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	sim := backends.NewSimulatedBackend(alloc, 1000)

	hash := common.HexToHash("")
	scAddress := common.BytesToAddress(hash[:])
	simpleLoan, _ := NewSimpleLoan(scAddress, sim)
	opts := &bind.TransactOpts{

	}
	simpleLoan.AcceptLoan(opts, [32]byte{}, [32]byte{}, [32]byte{})
}
