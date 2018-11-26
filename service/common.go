package service

import (
	"time"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
)

// Call blockchain to get tx in block by hash
// if can not get anything, we will retry after sleeping time
func GetBlockchainTxByHash(txId string, retry int, bc *blockchain.Blockchain) (*blockchain.TransactionDetail, error) {
	var tx *blockchain.TransactionDetail
	for true {
		var err error
		tx, err = bc.GetTxByHash(txId)
		if err != nil {
			return nil, err
		}
		// retry 10 times = 30s
		time.Sleep((3 * time.Millisecond))
		retry --
		if retry == 0 {
			break
		}
	}
	return tx, nil
}
