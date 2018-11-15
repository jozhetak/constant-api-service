package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/api"
	"github.com/ninjadotorg/constant-api-service/conf"
	"github.com/ninjadotorg/constant-api-service/dao"
	"github.com/ninjadotorg/constant-api-service/pubsub"
	"github.com/ninjadotorg/constant-api-service/service"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
	"github.com/ninjadotorg/constant-api-service/database"
	"github.com/ninjadotorg/constant-api-service/service/3rd/sendgrid"
	"github.com/ninjadotorg/constant-api-service/dao/portal"
	"github.com/ninjadotorg/constant-api-service/dao/exchange"
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

	mailClient := sendgrid.Init(conf)

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

		userDAO = dao.NewUser(db)
		userSvc = service.NewUserService(userDAO, bc)

		portalDAO = portal.NewPortal(db)
		portalSvc = service.NewPortal(portalDAO, bc)

		exchangeDAO = exchange.NewExchange(db)
		exchangeSvc = service.NewExchange(exchangeDAO)

		walletSvc = service.NewWallet(bc)

		pubsubSvc = pubsub.NewService()
	)

	r := gin.Default()
	svr := api.NewServer(r, pubsubSvc, upgrader, userSvc, portalSvc, exchangeSvc, walletSvc, logger, mailClient)
	authMw := api.AuthMiddleware(string(conf.TokenSecretKey), svr.Authenticate)
	svr.Routes(authMw)

	if err := r.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		logger.Fatal("router.Run", zap.Error(err))
	}
}
