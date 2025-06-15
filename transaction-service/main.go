// finbank/transaction-service/main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"transaction-service/handlers"
	"transaction-service/queue"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:password@postgres:5432/finbank?sslmode=disable")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	handlers.Init(db)
	queue.InitKafkaProducer()
	go queue.ProcessQueue()

	http.HandleFunc("/transaction", handlers.CreateTransactionHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Transaction Service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
