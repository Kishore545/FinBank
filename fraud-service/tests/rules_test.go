// finbank/fraud-service/tests/rules_test.go
package tests

import (
	"testing"
	"fraud-service/rules"
)

func TestIsFraudulent(t *testing.T) {
	rules.LoadRules()

	tx := rules.Transaction{
		ID:        1,
		UserID:    123,
		Amount:    15000,
		Timestamp: "2024-06-15T12:00:00Z",
	}

	if !rules.IsFraudulent(tx) {
		t.Error("Expected transaction to be flagged as fraudulent")
	}
}
