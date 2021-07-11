package messaging

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lfrei/kafka-store/store"
	"sync"
)

func processMessage(msg *kafka.Message) {
	id := string(msg.Key)
	product := string(msg.Value)
	store.AddProduct(id, product)
}

func Start(wg *sync.WaitGroup, topic string) {
	fmt.Println("Start Kafka Consumer for topic", topic)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "kafka-store",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		panic(err)
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			processMessage(msg)
		}
	}

	c.Close()
	wg.Done()
}
