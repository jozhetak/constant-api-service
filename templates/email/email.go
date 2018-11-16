package email

import sendgrid "github.com/sendgrid/sendgrid-go"

type Email struct {
	mailer *sendgrid.Client
}

func New(mailer *sendgrid.Client) *Email {
	return &Email{mailer}
}
