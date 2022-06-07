package utils

import (
	"fmt"
	"net/smtp"
	"rabbitmq/config"
)

func SendEmail(email string) {
	to := []string{email}

	address := config.Getconfig().SmtpHost + ":" + config.Getconfig().SmtpPort
	//build msg
	subject := "Rabbitmq"
	body := "sending email from rabbitmq"

	message := fmt.Sprintf("From: %s\r\n", config.Getconfig().Email)
	message += fmt.Sprintf("To: %s\r\n", to)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += fmt.Sprintf("\r\n%s\r\n", body)

	fmt.Println(message)
	auth := smtp.PlainAuth("", config.Getconfig().Username, config.Getconfig().SmtpPassword, config.Getconfig().SmtpHost)

	// send mail
	err := smtp.SendMail(address, auth, config.Getconfig().Email, to, []byte(message))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("email processed")

}
