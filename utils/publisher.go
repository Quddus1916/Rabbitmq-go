package utils

import (
	"fmt"
	"github.com/streadway/amqp"
)

func Publish(email string) error {
	//connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()
	if err != nil {
		return err
	}
	fmt.Println("connected")
	//channel
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}
	//queue
	q, err := ch.QueueDeclare(
		"email",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}
	//publish
	err = ch.Publish(

		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(email),
		})
	fmt.Println("proccessed")
	return err
}
