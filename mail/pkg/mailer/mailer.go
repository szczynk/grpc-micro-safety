package mailer

import (
	"bytes"
	"crypto/tls"
	"errors"
	"mail/config"
	netmail "net/mail"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/k3a/html2text"
	mail "github.com/xhit/go-simple-mail/v2"
)

type EmailData struct {
	URL      string
	Email    string
	Username string
	Subject  string
}

type Mail interface {
	SendEmail(data *EmailData, emailTemplate string) error
}

type Mailer struct {
	mailer *mail.SMTPServer
	from   string
}

func NewMailer(cfg *config.Config) (*Mailer, error) {
	port, err := strconv.Atoi(cfg.SMTP.Port)
	if err != nil {
		return nil, err
	}
	if port == 0 {
		return nil, errors.New("SMTP Port is 0")
	}

	from := netmail.Address{
		Name:    cfg.SMTP.SenderIdentity,
		Address: cfg.SMTP.SenderEmail,
	}

	mailer := mail.NewSMTPClient()
	mailer.Host = cfg.SMTP.Host
	mailer.Port = port
	mailer.Username = cfg.SMTP.User
	mailer.Password = cfg.SMTP.Pass
	mailer.Encryption = mail.EncryptionSTARTTLS
	// auth = AuthPlain, AuthLogin, AuthCRAMMD5, and AuthNone. Default is AuthPlain

	if cfg.Server.Mode == "development" {
		mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return &Mailer{
		mailer: mailer,
		from:   from.String(),
	}, nil
}

func (m *Mailer) SendEmail(data *EmailData, emailTemplate string) error {
	template, err := parseTemplateDir("internal/templates")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = template.ExecuteTemplate(&body, emailTemplate, &data)
	if err != nil {
		return err
	}

	smtpClient, err := m.mailer.Connect()
	if err != nil {
		return err
	}

	// New email simple html with inline and CC
	email := mail.NewMSG()
	email.SetFrom(m.from).
		AddTo(data.Email).
		SetSubject(data.Subject)

	email.GetFrom()
	email.SetBody(mail.TextHTML, body.String())
	email.AddAlternative(mail.TextPlain, html2text.HTML2Text(body.String()))

	if email.Error != nil {
		return email.Error
	}

	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	return nil
}

// 👇 Email template parser
func parseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
