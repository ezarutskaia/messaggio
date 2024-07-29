package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

func MessageProducer(message string) {
	topic := "messages"

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"acks": "all"})
	
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: []byte(message)},
		nil,
	)

	if err != nil {
		fmt.Printf("Failed to send message: %s\n", err)
		os.Exit(1)
	}

	p.Flush(1000)
}