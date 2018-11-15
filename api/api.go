package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/pubsub"
	"github.com/ninjadotorg/constant-api-service/service"
)

type Server struct {
	g           *gin.Engine
	up          *websocket.Upgrader
	ps          *pubsub.Service
	mailer      *sendgrid.Client
	userSvc     *service.User
	portalSvc   *service.Portal
	exchangeSvc *service.Exchange
	walletSvc   *service.Wallet
	logger      *zap.Logger
}

func NewServer(g *gin.Engine, ps *pubsub.Service, up *websocket.Upgrader, userSvc *service.User, portalSvc *service.Portal, exchangeSvc *service.Exchange, walletSvc *service.Wallet, logger *zap.Logger, mailer *sendgrid.Client) *Server {
	return &Server{
		g:           g,
		up:          up,
		ps:          ps,
		mailer:      mailer,
		userSvc:     userSvc,
		portalSvc:   portalSvc,
		exchangeSvc: exchangeSvc,
		walletSvc:   walletSvc,
		logger:      logger,
	}
}

type response struct {
	Data  interface{}    `json:"Data"`
	Error *service.Error `json:"Error"`
}
