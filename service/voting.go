package service

import (
	"github.com/ninjadotorg/constant-api-service/dao/voting"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
)

type VotingService struct {
	votingDao         *voting.VotingDao
	blockchainService *blockchain.Blockchain
}

func NewVotingService(r *voting.VotingDao, bc *blockchain.Blockchain) *VotingService {
	return &VotingService{
		votingDao:         r,
		blockchainService: bc,
	}
}
