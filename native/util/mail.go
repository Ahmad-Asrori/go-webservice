package util

import (
	"fmt"
	"net/smtp"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_EMAIL = "EMAIL"
const CONFIG_PASSWORD = "PASSWORD"

func SendMail(to []string, username, subject, message string) error {
	body := message

	auth := smtp.PlainAuth("", CONFIG_EMAIL, CONFIG_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_EMAIL, to, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
