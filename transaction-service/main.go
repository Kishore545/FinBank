package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
)

type Transaction struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

var kafkaWriter *kafka.Writer

func main() {
	// Kafka writer
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"), // change to "localhost:9092" if running outside Docker
		Topic:    "transactions",
		Balancer: &kafka.LeastBytes{},
	}

	// Endpoints
	http.HandleFunc("/transactions", handleTransaction)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	fmt.Println("🚀 Transaction service running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var tx Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Add timestamp if not set
	if tx.Timestamp == "" {
		tx.Timestamp = time.Now().UTC().Format(time.RFC3339)
	}

	// Send to Kafka
	msg, _ := json.Marshal(tx)
	err := kafkaWriter.WriteMessages(r.Context(),
		kafka.Message{Value: msg},
	)
	if err != nil {
		log.Println("Kafka write error:", err)
		http.Error(w, "Failed to send transaction", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Transaction sent: %+v\n", tx)
	w.WriteHeader(http.StatusAccepted)
}
