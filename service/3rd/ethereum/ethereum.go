package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthClient struct {
	Conn    *ethclient.Client
	Context context.Context
}

func CreateEthClient() *EthClient {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal("Whoops something went wrong!", err)
	}

	ctx := context.Background()
	return &EthClient{Conn: conn, Context: ctx}
}

func (self EthClient) TransactionByHash(hash string) (tx *types.Transaction, isPending bool, err error) {
	tx, pending, err := self.Conn.TransactionByHash(self.Context, common.HexToHash(hash))
	if !pending {
		fmt.Println(tx)
	}
	return tx, pending, err
}

func (self EthClient) CallConstract() error {
	return nil
}
