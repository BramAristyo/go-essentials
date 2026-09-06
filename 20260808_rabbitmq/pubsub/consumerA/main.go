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

	err = ch.Qos(
		1,
		0,
		false,
	)

	if err != nil {
		log.Fatal(err)
	}

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

	_, err = ch.QueueDeclare(
		"log-queue",
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	// subscribe to 'logs' exchange
	err = ch.QueueBind(
		"log-queue",
		"",
		"logs",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		"log-queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			time.Sleep(1 * time.Second)
			log.Printf("A Received: %s", msg.Body)
			log.Print("Done")
			msg.Ack(false) // if operation success it will be drop the queue from the list
		}
	}()

	<-forever
}
