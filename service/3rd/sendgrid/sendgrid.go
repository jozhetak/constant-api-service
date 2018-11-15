package sendgrid

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/ninjadotorg/constant-api-service/conf"
)

func Init(conf *config.Config) (*sendgrid.Client) {
	mailClient := sendgrid.NewSendClient(conf.SendgridAPIKey)
	return mailClient
}
