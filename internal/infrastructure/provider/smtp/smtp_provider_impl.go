package smtp

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/InstayPMS/backend/internal/application/port"
	"github.com/InstayPMS/backend/internal/infrastructure/config"
)

//go:embed templates/auth.html
var authTemplate embed.FS

type AuthEmailData struct {
	Subject string `json:"subject"`
	Otp     string `json:"otp"`
}

type smtpProviderImpl struct {
	cfg  config.SMTPConfig
	auth smtp.Auth
}

func NewSMTPProvider(cfg config.SMTPConfig) port.SMTPProvider {
	auth := smtp.PlainAuth("", cfg.User, cfg.Password, cfg.Host)
	return &smtpProviderImpl{
		cfg,
		auth,
	}
}

func (s *smtpProviderImpl) Send(to, subject, body string) error {
	msg := fmt.Appendf(nil, "Subject: %s\r\nMIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n%s", subject, body)
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	return smtp.SendMail(addr, s.auth, s.cfg.User, []string{to}, msg)
}

func (s *smtpProviderImpl) AuthEmail(to, subject, otp string) error {
	tmpl, err := template.ParseFS(authTemplate, "templates/auth.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	data := AuthEmailData{
		Subject: subject,
		Otp:     otp,
	}
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	return s.Send(to, subject, body.String())
}
