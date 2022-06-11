package utils

import (
	"sister-backend/config"

	"github.com/go-gomail/gomail"
)

type MailMessage struct {
	*gomail.Message
}

func NewMailConnection(cfg config.MailConfig) *gomail.Dialer {
	return gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
}

func SendEmails(conn *gomail.Dialer, messages ...MailMessage) error {
	for _, msg := range messages {
		if err := conn.DialAndSend(msg.Message); err != nil {
			return err
		}
	}
	return nil
}

func SendEmail(conn *gomail.Dialer, message MailMessage) error {
	if err := conn.DialAndSend(message.Message); err != nil {
		return err
	}
	return nil
}

func NewMailMessage() (msg *MailMessage) {
	msg = &MailMessage{
		Message: gomail.NewMessage(),
	}

	msg.SetHeader("From", "admin@toqcer.id")

	return msg
}

func (msg MailMessage) SetReceiver(receivers ...string) {
	msg.Message.SetHeader("To", receivers...)
}

func (msg MailMessage) SetSubject(subject string) {
	msg.Message.SetHeader("Subject", subject)
}

func (msg MailMessage) SetBody(body string) {
	msg.Message.SetBody("text/html", body)
}

func (msg MailMessage) AddAttachment(file string) {
	msg.Message.Attach(file)
}

func (msg MailMessage) AddAttachments(files ...string) {
	for _, file := range files {
		msg.AddAttachment(file)
	}
}
