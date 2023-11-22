package adapter

import (
	"fww-wrapper/internal/container/infrastructure/email"
	"fww-wrapper/internal/data/dto_notification"
)

// SendEmailNotification implements Adapter.
func (a *adapter) SendEmailNotification(data *dto_notification.SendEmailRequest) (err error) {
	data.EmailAddress = a.emailConfig.EmailAddress
	// Init Email SMTP
	dial, mail := email.InitializeSendEmail(a.emailConfig)

	// TODO: Populate data body base on route

	// compose email
	mail = email.ComposeEmail(mail, data, nil)
	// Send E-Mail
	if err := dial.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
