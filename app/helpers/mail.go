package helpers

import (
	"crypto/tls"
	"os"
	"strconv"
	"strings"

	gomail "gopkg.in/mail.v2"
)

type IMailServer interface {
	Send() error
	SendCopyToAdmin() error
}

type MailServer struct {
	From           string
	To             string
	Subject        string
	Host           string
	Port           int
	Username       string
	Password       string
	Body           string
	SuperAdminMail string
}

var mailServer MailServer

func NewMailServer(to []string, subject string) MailServer {
	mailServer.From = os.Getenv("MAIL_FROM_NAME")
	mailServer.To = strings.Join(to, ",")
	mailServer.Subject = subject
	mailServer.Host = os.Getenv("MAIL_HOST")
	mailServer.Port, _ = strconv.Atoi(os.Getenv("MAIL_PORT"))
	mailServer.Username = os.Getenv("MAIL_USERNAME")
	mailServer.Password = os.Getenv("MAIL_PASSWORD")
	mailServer.SuperAdminMail = os.Getenv("MAIL_SUPER_ADMIN")
	return mailServer
}

func (s *MailServer) Send() error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", s.From)
	mail.SetHeader("To", s.To)
	mail.SetHeader("Subject", s.Subject)

	mail.SetBody("text/html", s.Body)

	d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(mail)
}

func (s *MailServer) SendCopyToAdmin() error {
	s.To = s.SuperAdminMail
	return s.Send()
}
