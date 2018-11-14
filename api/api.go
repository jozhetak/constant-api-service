package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/service"
)

type Server struct {
	g           *gin.Engine
	up          *websocket.Upgrader
	hub         *hub
	mailer      *sendgrid.Client
	userSvc     *service.User
	portalSvc   *service.Portal
	exchangeSvc *service.Exchange
	logger      *zap.Logger
}

func NewServer(g *gin.Engine, up *websocket.Upgrader, userSvc *service.User, portalSvc *service.Portal, exchangeSvc *service.Exchange, logger *zap.Logger, mailer *sendgrid.Client) *Server {
	h := newHub()
	go h.run()

	return &Server{
		g:           g,
		up:          up,
		hub:         h,
		mailer:      mailer,
		userSvc:     userSvc,
		portalSvc:   portalSvc,
		exchangeSvc: exchangeSvc,
		logger:      logger,
	}
}

type response struct {
	Data  interface{}    `json:"data"`
	Error *service.Error `json:"error"`
}
