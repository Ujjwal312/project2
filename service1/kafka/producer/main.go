package main

import (
	"fmt"
	"log"
	

	"github.com/IBM/sarama"
)



type order struct {
	Id string
	Product string
	Userid string 
}

  [    {
        "id":"3",
        "product":"Laptop",
        "userid":"ujjwal@gmail.com"
      },

      {
        "id":"4",
        "product":"Laptop",
        "userid":"ujjwal@gmail.com"
      },
      {
        "id":"6",
        "product":"Laptop",
        "userid":"ujjwal@gmail.com"
      },
      {
        "id":"7",
        "product":"Laptop",
        "userid":"ujjwal@gmail.com"
      },
      {
        "id":"8",
        "product":"Laptop",
        "userid":"ujjwal@gmail.com"
      },
      {
        "id":"9",
        "product":"Laptop",
        "userid":"ujjwal@gmail.com"
      },
      {
        "id":"10",
        "product":"Laptop",
        "userid":"ujjwal@gmail.com"
      },
      {
        "id":"11",
        "product":"Laptop",
        "userid":"ujjwal@gmail.com"
      },
      {
        "id":"12",
        "product":"Laptop",
        "userid":"ujjwal@gmail.com"
      }
	
  ]

func main() {
a:=&order{} 
	a.Id ="20"
	a.Product ="Laptop"
	a.Userid ="ujjwal@gmail.com"
	broker := "localhost:9092" 
	topic := "order"  

	
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	
	
	
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder{a},
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		} else {
			log.Printf("Message sent: partition=%d, offset=%d", partition, offset)
		}
	

	fmt.Println("All messages sent successfully!")
}
