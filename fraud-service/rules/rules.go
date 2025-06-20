package rules

import "fraud/models"

type Transaction struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

// Optional for future rule loading; no-op for now
func LoadRules() {}

func IsFraudulent(tx models.Transaction) bool {
	// Rule: flag as fraudulent if amount > 10,000
	return tx.Amount > 100000
}
