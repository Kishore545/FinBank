package handlers

import (
	"encoding/json"
	"net/http"
	"transaction/models"
)

var db DB

func Init(database DB) {
	db = database
}

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var tx models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := db.InsertTransaction(tx); err != nil {
		http.Error(w, "DB insert error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
