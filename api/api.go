package api

import (
	"strconv"

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
	exchangeSvc *service.ExchangeService
	walletSvc   *service.WalletService
	votingSvc   *service.VotingService
	reserveSvc  *service.ReserveService
	logger      *zap.Logger
}

func (s *Server) pagingFromContext(c *gin.Context) (int, int) {
	var (
		pageS  = c.DefaultQuery("page", "1")
		limitS = c.DefaultQuery("limit", "10")
		page   int
		limit  int
		err    error
	)

	page, err = strconv.Atoi(pageS)
	if err != nil {
		page = 1
	}

	limit, err = strconv.Atoi(limitS)
	if err != nil {
		limit = 10
	}

	return page, limit
}

func NewServer(g *gin.Engine, ps *pubsub.Pubsub, up *websocket.Upgrader, userSvc *service.User, portalSvc *service.PortalService, votingSvc *service.VotingService, exchangeSvc *service.ExchangeService, walletSvc *service.WalletService, reserveSvc *service.ReserveService, logger *zap.Logger) *Server {
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
		reserveSvc:  reserveSvc,
	}
}
