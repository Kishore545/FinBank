// finbank/transaction-service/handlers/transaction.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"transaction-service/models"
	"transaction-service/queue"
)

var db *sql.DB

func Init(database *sql.DB) {
	db = database
}

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var tx models.Transaction
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = models.CreateTransaction(db, &tx)
	if err != nil {
		http.Error(w, "Failed to store transaction", http.StatusInternalServerError)
		return
	}

	queue.Enqueue(tx)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Transaction received"))
}
