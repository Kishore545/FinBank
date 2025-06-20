package tests

import (
	"fraud/models"
	"fraud/rules"
	"testing"
)

func TestIsFraudulent(t *testing.T) {
	rules.LoadRules()

	tx := models.Transaction{ // ✅ use models.Transaction
		ID:        1,
		UserID:    123,
		Amount:    15000,
		Timestamp: "2024-06-15T12:00:00Z",
	}

	if !rules.IsFraudulent(tx) {
		t.Error("Expected transaction to be flagged as fraudulent")
	}
}
