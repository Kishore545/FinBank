package kafka

import (
	"context"
	"encoding/json"
	"os"
	"time"
	"transaction/models"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func init() {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "localhost:9092"
	}
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    "transactions",
		Balancer: &kafka.LeastBytes{},
	})
}

func Publish(tx models.Transaction) error {
	data, err := json.Marshal(tx)
	if err != nil {
		return err
	}
	return writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(string(tx.UserID)),
			Value: data,
			Time:  time.Now(),
		})
}
