// finbank/fraud-service/rules/engine.go
package rules

type Transaction struct {
	ID        int
	UserID    int
	Amount    float64
	Timestamp string
}

var threshold float64

func LoadRules() {
	threshold = 10000 // Static rule: flag large transactions
}

func IsFraudulent(tx Transaction) bool {
	return tx.Amount > threshold
}
