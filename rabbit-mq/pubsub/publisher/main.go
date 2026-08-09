package main

import (
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// declare new exchange with 'fanout' kind
	err = ch.ExchangeDeclare(
		"logs",
		amqp091.ExchangeFanout,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for range 5 {
		// publish message to 'logs' exchange
		err = ch.PublishWithContext(
			ctx,
			"logs",
			"",
			true,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte("Hello from pubsub publisher"),
			},
		)

		if err != nil {
			log.Fatal(err)
		}
	}
}
