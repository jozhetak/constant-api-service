package service

import (
	"github.com/ninjadotorg/constant-api-service/dao/reserve"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
)

type ReserveService struct {
	reserveDao        *reserve.ReserveDao
	blockchainService *blockchain.Blockchain
}

func NewReserveService(reserveDao *reserve.ReserveDao, blockchainService *blockchain.Blockchain) *ReserveService {
	return &ReserveService{
		reserveDao:        reserveDao,
		blockchainService: blockchainService,
	}
}
