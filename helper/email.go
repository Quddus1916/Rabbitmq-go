package helper

import (
	//"net/smtp"
	"rabbitmq/utils"
)

func PublishEmail(email string) error {
	err := utils.Publish(email)

	return err
}
