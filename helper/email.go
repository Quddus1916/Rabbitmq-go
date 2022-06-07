package helper

import (
	//"net/smtp"
	"rabbitmq/utils"
)

func SendEmail(email string) error {
	err := utils.Publish(email)

	return err
}
