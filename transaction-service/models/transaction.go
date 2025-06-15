// finbank/transaction-service/models/transaction.go
package models

import (
	"database/sql"
	"fmt"
)

type Transaction struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

func CreateTransaction(db *sql.DB, tx *Transaction) error {
	query := `INSERT INTO transactions (user_id, amount, timestamp) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, tx.UserID, tx.Amount, tx.Timestamp)
	return err
}
