// finbank/fraud-service/kafka/consumer.go
package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"fraud-service/rules"
	"fraud-service/alerts"
)

type Transaction struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

func ConsumeTransactions() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "transactions",
		GroupID: "fraud-detector",
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading Kafka message:", err)
			continue
		}

		var tx Transaction
		err = json.Unmarshal(m.Value, &tx)
		if err != nil {
			log.Println("Invalid transaction format:", err)
			continue
		}

		if rules.IsFraudulent(tx) {
			alerts.SendAlert(tx)
		}
	}
}
