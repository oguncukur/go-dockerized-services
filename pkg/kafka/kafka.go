package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

func Producer() {

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	topic := "my-topic"
	message := "Hello, Kafkaaaaa!"

	err = producer.Produce(&kafka.Message{
		Key:            []byte(""),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		log.Fatalf("Error producing message: %v", err)
	}

	log.Println("Message sent successfully!")
}

func Consumer() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "test-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	topic := "test"
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %s\n", string(msg.Value))
		} else {
			log.Printf("Error reading message: %v\n", err)
		}
	}
}
