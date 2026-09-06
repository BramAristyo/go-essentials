package main

import (
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

	// configuration prefetch for consumer
	// only consume (n) message per operation
	err = ch.Qos(
		1,
		0,
		false,
	)

	// add more queue declare for type safety
	// recover error queue not found!
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

	// consume message from 'hello-queue' queue
	msgs, err := ch.Consume(
		"hello-queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			time.Sleep(3 * time.Second)
			log.Printf("Received: %s", msg.Body)
			log.Print("Done")
		}
	}()

	<-forever
}
