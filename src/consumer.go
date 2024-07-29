package main

import (
	"fmt"
	"log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"src/db"
)

func main() {
	fmt.Println("Consumer started.")
	topic := []string{"messages"}
	
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id": "foo",
		"go.application.rebalance.enable": true})
	
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}
	
	err = consumer.SubscribeTopics(topic, nil)

	run := true
	for run == true {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				consumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				consumer.Unassign()
			case *kafka.Message:
				fmt.Printf("Message: %s\n", string(e.Value))
				engine := *db.Engine()
				var message db.Messages
				err = engine.First(&message, string(e.Value)).Error
				if err == nil {									
					message.Processed = true
					err = engine.Save(&message).Error
					if err != nil {
						log.Fatalf("failed to update user: %v", err)
					}
				}
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
		}
	}
	
	consumer.Close()
}