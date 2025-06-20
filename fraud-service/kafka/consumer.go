package kafka

import (
	"context"
	"encoding/json"
	"log"

	"fraud/alerts"
	"fraud/models"
	"fraud/rules"

	"github.com/segmentio/kafka-go"
)

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

		var tx models.Transaction
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
