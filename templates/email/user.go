package email

import (
	"bytes"
	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/pkg/errors"
	sendgridMail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func (e *Email) SendForgotPasswordEmail(data *ForgotPassword) error {
	html, text, err := getForgotPasswordBody(data)
	if err != nil {
		return errors.Wrap(err, "getForgotPasswordBody")
	}

	var (
		from    = sendgridMail.NewEmail("Constant", "human@autonomous.ai")
		to      = sendgridMail.NewEmail(data.Name, data.Email)
		message = sendgridMail.NewSingleEmail(from, "Reset your password", to, html, text)
	)

	_, err = e.mailer.Send(message)
	if err != nil {
		return errors.Wrap(err, "u.mailer.Send")
	}
	return nil
}

type ForgotPassword struct {
	Email string
	Name  string
	Link  string
}

const forgotPasswordText = `
Hi {{.Name}},

Please click this link {{.Link}} to reset your password.
`

const forgotPasswordHTML = `
Hi {{.Name}}, <br /><br />

Please click this link {{.Link}} to reset your password.
`

func getForgotPasswordTemplate() (html *htmlTemplate.Template, text *textTemplate.Template, err error) {
	html, err = htmlTemplate.New("html").Parse(forgotPasswordHTML)
	if err != nil {
		return nil, nil, errors.Wrap(err, "t.Parse")
	}
	text, err = textTemplate.New("text").Parse(forgotPasswordText)
	if err != nil {
		return nil, nil, errors.Wrap(err, "t.Parse")
	}
	return
}

func getForgotPasswordBody(f *ForgotPassword) (html, text string, err error) {
	h, t, err := getForgotPasswordTemplate()
	if err != nil {
		return "", "", errors.Wrap(err, "email.GetForgotPasswordTemplate")
	}

	var htmlBody bytes.Buffer
	if err := h.Execute(&htmlBody, f); err != nil {
		return "", "", errors.Wrap(err, "html.Execute")
	}

	var textBody bytes.Buffer
	if err := t.Execute(&textBody, f); err != nil {
		return "", "", errors.Wrap(err, "text.Execute")
	}

	return htmlBody.String(), textBody.String(), nil
}
