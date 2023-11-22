package email

import (
	"bytes"
	"fmt"
	"fww-wrapper/internal/data/dto_notification"
	"io"
	"net/http"
	"strings"

	"gopkg.in/gomail.v2"
)

func ComposeEmail(m *gomail.Message, data *dto_notification.SendEmailRequest, attachmentObjects []io.Reader) *gomail.Message {
	// set Header for email
	to := data.To

	splittedTo := strings.Split(data.To, "|")
	if len(splittedTo) > 1 {
		to = splittedTo[1]
	}

	m.SetHeader("From", data.EmailAddress)
	m.SetHeader("To", to)
	if data.Cc != "" {
		m.SetHeader("Cc", data.Cc)
	}
	if data.Bcc != "" {
		m.SetHeader("Bcc", data.Bcc)
	}
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", data.Body)
	for i, attachmentObject := range attachmentObjects {
		// objectBuffer := make([]byte, 0)

		objectBuffer := new(bytes.Buffer)
		_, err := objectBuffer.ReadFrom(attachmentObject)

		if err != nil {
			fmt.Println(err.Error())
		}

		objectBytes := objectBuffer.Bytes()
		// attachmentObject.Read(objectBuffer)

		contentType := http.DetectContentType(objectBytes)

		m.Attach(data.Attachments[i], gomail.SetCopyFunc(
			func(w io.Writer) error {
				_, err := w.Write(objectBytes)
				fmt.Println(err)
				return err
			}), gomail.SetHeader(map[string][]string{
			"Content-Type": {
				contentType,
			}}))

	}

	return m
}
