package email_user

import (
	"fmt"
	"net/smtp"
	"os"
	"project-x/internal/utils"
)

func MailUser(reciever, subject string) {
	from := os.Getenv("SUPPORT_EMAIL")
	password := os.Getenv("PASSWORD")

	to := []string{
		reciever,
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	msg := `To: ` + reciever + "\n" +
		`From: ` + from + "\n" +
		`Subject: Testing from GoLang` + "\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		`<html><body><h1>Hello World!</h1></body></html>`

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}
	utils.Logger.Info("Mail executed")
}
