package util

import (
	"NGB/pkg/logrus"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type CustomEmail struct {
	From    string
	To      string
	Subject string
	Content string
	Account string
	Code    string
	Addr    string
	Server  string
}

func SendEmail(customEmail CustomEmail) error {
	logrus.Logger.Infof("Send email: %v", customEmail)
	e := email.NewEmail()
	e.From = customEmail.From
	e.To = []string{customEmail.To}
	e.Subject = customEmail.Subject
	e.HTML = []byte(customEmail.Content)
	return e.Send(customEmail.Addr, smtp.PlainAuth("", customEmail.Account, customEmail.Code, customEmail.Server))
}
