package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/utils"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("sending mail to: %s", d.Body)
			utils.SendEmail(string(d.Body))

		}
	}()

	log.Printf(" [*] Waiting for emails. To exit press CTRL+C")
	<-forever

}
