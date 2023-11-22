package controller

import (
	"fww-wrapper/internal/data/dto_notification"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
)

func (c *Controller) SendEmailNotificationHandler(msg *message.Message) error {
	var body dto_notification.SendEmailRequest

	if err := json.Unmarshal(msg.Payload, &body); err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}

	err := c.Adapter.SendEmailNotification(&body)
	if err != nil {
		msg.Ack()
		c.Log.Error(err)
		return err
	}

	msg.Ack()
	return nil
}
