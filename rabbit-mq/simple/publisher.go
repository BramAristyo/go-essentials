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

	// add new queue
	_, err = ch.QueueDeclare(
		"hello-queue",
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// publish message to 'hello-queue' queue from default RabbitMQ exchange ""
	// with body text/plain
	err = ch.PublishWithContext(ctx,
		"",
		"hello-queue",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello from publisher!"),
		},
	)

	if err != nil {
		log.Fatal(err)
	}
}
