package service

import (
	"github.com/ninjadotorg/constant-api-service/dao/reserve"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/ninjadotorg/constant-api-service/serializers"
	"github.com/ninjadotorg/constant-api-service/models"
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

func (self *ReserveService) RequestReserve(req *serializers.ReserveContributionRequest) {
	// 1. Validate ReserveContributionRequest in request
	// 2. insert db ReserveContributionRequest
	// 3. insert db ReserveContributionRequestPaymentParty
	// 4. call related party: prime trust, eth ...
	// 5. update data ReserveContributionRequestPaymentParty
	// 6. ... waitting webhook
}

func (self *ReserveService) GetRequestReserve(u *models.User) (*models.ReserveContributionRequest, error) {
	// TODO
	return nil, nil
}

func (self *ReserveService) RequestReserveHistory(u *models.User) ([]*models.ReserveContributionRequest, error) {
	// TODO
	return nil, nil
}

func (self *ReserveService) ReturnRequestReserve(req *serializers.ReserveDisbursementRequest) {
	// 1. Validate ReserveDisbursementRequest in request
	// 2. insert db ReserveDisbursementRequest
	// 3. insert db ReserveDisbursementRequestPaymentParty
	// 6. call blockchain network to burn constant
	// 4. call related party: prime trust, eth ... and wait for data
	// 5. update data ReserveContributionRequestPaymentParty
}

func (self *ReserveService) GetReturnRequestReserve(u *models.User) (*models.ReserveDisbursementRequest, error) {
	// TODO
	return nil, nil
}

func (self *ReserveService) ReturnRequestReserveHistory(u *models.User) ([]*models.ReserveDisbursementRequest, error) {
	// TODO
	return nil, nil
}

func (self *ReserveService) PrimetrustWebHook() {

}
