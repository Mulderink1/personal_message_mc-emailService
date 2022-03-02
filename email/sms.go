package email

import (
	"fmt"
	domain "github.com/Mulderink1/personal_message_mc-services-domain"
	"net/smtp"
	"os"
)

func SmsSender(log *domain.Logger, email string) {
	from := os.Getenv("EMAIL")
	pass := os.Getenv("PASSWORD")

	to := []string{
		email,
	}

	msg := []byte("This is a test email message from GoLang.")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", from, pass, smtpHost)

	smtpConn := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	err := smtp.SendMail(smtpConn, auth, from, to, msg)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	log.Logger.Info(fmt.Sprintf("Email Sent to: %v", email))
}
