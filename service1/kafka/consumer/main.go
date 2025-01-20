package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	broker := "localhost:9092" 
	topic := "order"  

	
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("Error fetching partitions: %v", err)
	}

	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Error starting partition consumer: %v", err)
		}
		defer pc.Close()

		
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				
			}
		}(pc)
	}

	
	select {}
}
