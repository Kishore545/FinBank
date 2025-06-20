package queue

import (
	"context"
	"encoding/json"
	"log"

	"transaction/models"

	"github.com/segmentio/kafka-go"
)

var (
	producer *kafka.Writer
	txQueue  = make(chan models.Transaction, 100)
)

func InitKafkaProducer() {
	producer = &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "transactions",
		Balancer: &kafka.LeastBytes{},
	}
}

func Enqueue(tx models.Transaction) {
	txQueue <- tx
}

func ProcessQueue() {
	for tx := range txQueue {
		msg, err := json.Marshal(tx)
		if err != nil {
			log.Println("Failed to marshal transaction:", err)
			continue
		}

		err = producer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(string(tx.ID)),
			Value: msg,
		})
		if err != nil {
			log.Println("Failed to send to Kafka:", err)
		}
	}
}
