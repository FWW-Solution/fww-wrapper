package email

import (
	"crypto/tls"
	"fww-wrapper/internal/config"
	"strconv"

	"gopkg.in/gomail.v2"
)

func InitializeSendEmail(cfg *config.EmailConfig) (d *gomail.Dialer, m *gomail.Message) {
	port, err := strconv.Atoi(cfg.SmtpPort)
	if err != nil {
		panic(err)
	}
	d = &gomail.Dialer{Host: cfg.Server, Port: port}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: cfg.SkipSSL}

	m = gomail.NewMessage()

	return

}
