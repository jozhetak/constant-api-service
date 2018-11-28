package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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
		bc     = blockchain.New(client, "http://127.0.0.1:9334")

		mailClient  = sendgrid.Init(conf)
		emailHelper = email.New(mailClient)

		userDAO = dao.NewUser(db)
		userSvc = service.NewUserService(userDAO, bc, emailHelper)

		portalDAO = portal.NewPortal(db)
		portalSvc = service.NewPortal(portalDAO, bc)

		votingDao = voting.NewVoting(db)
		votingSvc = service.NewVotingService(votingDao, bc)

		exchangeDAO = exchange.NewExchange(db)
		walletSvc   = service.NewWalletService(bc, exchangeDAO)
		exchangeSvc = service.NewExchange(exchangeDAO, walletSvc)
	)
	gcPubsubClient, err := gcloud.NewClient(context.Background(), "cash-prototype")
	if err != nil {
		logger.Fatal("gcloud.NewClient", zap.Error(err))
	}
	psSvc := pubsub.New(gcPubsubClient, exchangeDAO, bc, logger.With(zap.String("module", "pubsub")))

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))
	svr := api.NewServer(r, psSvc, upgrader, userSvc, portalSvc, votingSvc, exchangeSvc, walletSvc, logger)
	authMw := api.AuthMiddleware(string(conf.TokenSecretKey), svr.Authenticate)
	svr.Routes(authMw)

	if err := r.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		logger.Fatal("router.Run", zap.Error(err))
	}
}
