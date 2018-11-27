package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/pubsub"
	"github.com/ninjadotorg/constant-api-service/service"
)

type Server struct {
	g           *gin.Engine
	up          *websocket.Upgrader
	ps          *pubsub.Pubsub
	userSvc     *service.User
	portalSvc   *service.PortalService
	exchangeSvc *service.Exchange
	walletSvc   *service.WalletService
	votingSvc   *service.VotingService
	logger      *zap.Logger
}

func NewServer(g *gin.Engine, ps *pubsub.Pubsub, up *websocket.Upgrader, userSvc *service.User, portalSvc *service.PortalService, votingSvc *service.VotingService, exchangeSvc *service.Exchange, walletSvc *service.WalletService, logger *zap.Logger) *Server {
	return &Server{
		g:           g,
		up:          up,
		ps:          ps,
		userSvc:     userSvc,
		portalSvc:   portalSvc,
		exchangeSvc: exchangeSvc,
		walletSvc:   walletSvc,
		logger:      logger,
		votingSvc:   votingSvc,
	}
}
