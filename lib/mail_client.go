package lib

import (
	"fmt"
	"net/smtp"

	"example.com/config"
)

type MailClient struct {
	Config *config.SmtpConfig
}

func (s *MailClient) SendMail(to []string, subject string, body string) (err error) {
	message := fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n\r\n"+
			"%s\r\n",
		s.Config.From, to, subject, body,
	)

	addr := fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)

	auth := smtp.PlainAuth("", s.Config.Username, s.Config.Password, s.Config.Host)

	err = smtp.SendMail(addr, auth, s.Config.From, to, []byte(message))
	return
}
