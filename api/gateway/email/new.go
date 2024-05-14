package email

import (
	"bytes"
	"fmt"
	"mime"
	"net/mail"
	"net/smtp"
	"time"

	"github.com/jaytaylor/html2text"
	gomail "gopkg.in/mail.v2"

	"PopcornMovie/config"
)

type MailSender interface {
	SendMail(smtpTo string, subject, htmlBody string) error
}

type smtpConnectionInfo struct {
	SmtpUsername   string
	SmtpPassword   string
	SmtpServerHost string
	SmtpPort       string
	SenderEmail    string
}

func New(appConfig config.Configurations) MailSender {
	smtpConfig := appConfig.SMTP
	return &smtpConnectionInfo{
		SmtpServerHost: smtpConfig.SMTPHost,
		SmtpPort:       smtpConfig.SMTPPort,
		SmtpUsername:   smtpConfig.SMTPUsername,
		SmtpPassword:   smtpConfig.SMTPPassword,
		SenderEmail:    smtpConfig.SenderEmail,
	}
}

func (s smtpConnectionInfo) sendMail(from string, to []string, msg []byte) error {
	smtpAddress := s.SmtpServerHost + ":" + s.SmtpPort
	auth := smtp.PlainAuth("", s.SmtpUsername, s.SmtpPassword, s.SmtpServerHost)
	return smtp.SendMail(smtpAddress, auth, from, to, msg)
}

func (s smtpConnectionInfo) SendMail(smtpTo string, subject, htmlBody string) error {
	from := mail.Address{Address: s.SenderEmail}
	htmlMessage := "\r\n<html><body>" + htmlBody + "</body></html>"
	txtBody, err := html2text.FromString(htmlBody)
	if err != nil {
		txtBody = ""
	}
	headers := map[string][]string{
		"From":         {from.String()},
		"To":           {smtpTo},
		"Subject":      {encodeRFC2047Word(subject)},
		"MIME-Version": {"1.0"},
	}
	m := gomail.NewMessage(gomail.SetCharset("UTF-8"))
	m.SetHeaders(headers)
	m.SetDateHeader("Date", time.Now())
	m.SetBody("text/plain", txtBody)
	m.AddAlternative("text/html", htmlMessage)
	// Convert gomail.Message to []byte
	var buf bytes.Buffer
	if _, err := m.WriteTo(&buf); err != nil {
		fmt.Println("Error writing message to buffer:", err)
		return err
	}
	messageBytes := buf.Bytes()
	return s.sendMail(s.SenderEmail, []string{smtpTo}, messageBytes)
}
func encodeRFC2047Word(s string) string {
	return mime.BEncoding.Encode("utf-8", s)
}
