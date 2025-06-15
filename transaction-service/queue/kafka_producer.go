// finbank/transaction-service/queue/kafka_producer.go
package queue

import (
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"transaction-service/models"
)

var (
	producer *kafka.Writer
	txQueue = make(chan models.Transaction, 100)
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

		err = producer.WriteMessages(nil, kafka.Message{
			Key:   []byte(string(tx.ID)),
			Value: msg,
		})
		if err != nil {
			log.Println("Failed to send to Kafka:", err)
		}
	}
}
