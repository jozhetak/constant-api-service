package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"go.uber.org/zap"

	"github.com/ninjadotorg/constant-api-service/api"
	"github.com/ninjadotorg/constant-api-service/conf"
	"github.com/ninjadotorg/constant-api-service/dao"
	"github.com/ninjadotorg/constant-api-service/pubsub"
	"github.com/ninjadotorg/constant-api-service/service"
	"github.com/ninjadotorg/constant-api-service/service/3rd/blockchain"
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

	// pubsubClient, err := pubsub.NewClient(context.Background(), "test-pubsub-serv-1542155924980")
	// if err != nil {
	//         logger.Fatal("pubsub.NewClient", zap.Error(err))
	// }
	// ctx := context.Background()
	// topicName := "trades"
	// topic := pubsubClient.Topic(topicName)
	// exists, err := topic.Exists(ctx)
	// if err != nil {
	//         logger.Fatal("topic.Exists", zap.Error(err))
	// }
	// if !exists {
	//         logger.Info("creating topic", zap.String("name", topicName))
	//         _, err := pubsubClient.CreateTopic(ctx, "trades")
	//         if err != nil {
	//                 logger.Fatal("pubsubClient.CreateTopic", zap.Error(err))
	//         }
	// }

	mailer := sendgrid.NewSendClient(conf.SendgridAPIKey)

	db, err := gorm.Open("mysql", conf.Db)
	if err != nil {
		logger.Fatal("failed to open mysql db conn", zap.Error(err))
	}

	if err := dao.AutoMigrate(db); err != nil {
		logger.Fatal("failed to auto migrate", zap.Error(err))
	}

	var (
		client = &http.Client{}
		bc     = blockchain.New(client, "http://127.0.0.1:9334")

		userDAO = dao.NewUser(db)
		userSvc = service.NewUserService(userDAO, bc)

		portalDAO = dao.NewPortal(db)
		portalSvc = service.NewPortal(portalDAO, bc)

		exchangeDAO = dao.NewExchange(db)
		exchangeSvc = service.NewExchange(exchangeDAO)

		walletSvc = service.NewWallet(bc)

		pubsubSvc = pubsub.NewService()
	)

	r := gin.Default()
	svr := api.NewServer(r, pubsubSvc, upgrader, userSvc, portalSvc, exchangeSvc, walletSvc, logger, mailer)
	authMw := api.AuthMiddleware(string(conf.TokenSecretKey), svr.Authenticate)
	svr.Routes(authMw)

	if err := r.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		logger.Fatal("router.Run", zap.Error(err))
	}
}
