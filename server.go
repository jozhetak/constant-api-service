package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ninjadotorg/constant-api-service/service/3rd/ethereum"

	gcloud "cloud.google.com/go/pubsub"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/api"
	"github.com/ninjadotorg/constant-api-service/conf"
	"github.com/ninjadotorg/constant-api-service/dao"
	"github.com/ninjadotorg/constant-api-service/dao/exchange"
	"github.com/ninjadotorg/constant-api-service/dao/portal"
	"github.com/ninjadotorg/constant-api-service/dao/reserve"
	"github.com/ninjadotorg/constant-api-service/dao/voting"
	"github.com/ninjadotorg/constant-api-service/database"
	"github.com/ninjadotorg/constant-api-service/pubsub"
	"github.com/ninjadotorg/constant-api-service/service"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/ninjadotorg/constant-api-service/service/3rd/sendgrid"
	"github.com/ninjadotorg/constant-api-service/templates/email"
)

func main() {
	conf := config.GetConfig()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create zap logger: %v", err)
	}
	// defer logger.Sync()

	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	db, err := database.Init(conf)
	if err != nil {
		panic(err)
	}

	if err := dao.AutoMigrate(db); err != nil {
		logger.Fatal("failed to auto migrate", zap.Error(err))
	}

	var (
		client = &http.Client{}
		bc     = blockchain.New(client, conf.ConstantChainEndpoint)

		mailClient      = sendgrid.Init(conf)
		ethereumService = ethereum.Init(conf)
		emailHelper     = email.New(mailClient)

		userDAO = dao.NewUser(db)
		userSvc = service.NewUserService(userDAO, bc, emailHelper)

		portalDAO = portal.NewPortal(db)
		portalSvc = service.NewPortal(portalDAO, bc, ethereumService)

		votingDao = voting.NewVoting(db)
		votingSvc = service.NewVotingService(votingDao, bc)

		exchangeDAO = exchange.NewExchange(db)
		walletSvc   = service.NewWalletService(bc, exchangeDAO)
		exchangeSvc = service.NewExchange(exchangeDAO, walletSvc)

		reserveDAO = reserve.NewReserveDao(db)
		reserveSvc = service.NewReserveService(reserveDAO, bc)
	)
	gcPubsubClient, err := gcloud.NewClient(context.Background(), "cash-prototype")
	if err != nil {
		logger.Fatal("gcloud.NewClient", zap.Error(err))
	}
	psSvc := pubsub.New(gcPubsubClient, exchangeDAO, bc, logger.With(zap.String("module", "pubsub")), conf.OrderTopic, conf.OrderBookTopic, conf.OrderBookSubName)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))
	svr := api.NewServer(r, psSvc, upgrader, userSvc, portalSvc, votingSvc, exchangeSvc, walletSvc, reserveSvc, logger)
	authMw := api.AuthMiddleware(string(conf.TokenSecretKey), svr.Authenticate)
	svr.Routes(authMw)

	if err := r.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		logger.Fatal("router.Run", zap.Error(err))
	}
}
