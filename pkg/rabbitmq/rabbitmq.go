package rabbitmq

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func Publisher() {
	conn, err := amqp091.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatalf("RabbitMQ connection error: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Create channel error: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"test-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Create queue error: %s", err)
	}

	body := "Hello, RabbitMQ!"

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Publish message error: %s", err)
	}

	fmt.Printf("Sent message: %s\n", body)
}

func Consumer() {
	conn, err := amqp091.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatalf("RabbitMQ connection error: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Create channel error: %s", err)
	}
	defer ch.Close()

	messages, err := ch.Consume(
		"test-queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Consume message error: %s", err)
	}

	for message := range messages {
		fmt.Printf("Message is: %s\n", message.Body)
	}
}
