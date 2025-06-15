// finbank/transaction-service/tests/benchmark_test.go
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"transaction-service/handlers"
	"transaction-service/models"
)

func BenchmarkCreateTransactionHandler(b *testing.B) {
	db := setupMockDB() // assume this returns a working mock or real test DB
	handlers.Init(db)
	tx := models.Transaction{
		ID:        "tx-bench",
		UserID:    "user-bench",
		Amount:    123.45,
		Timestamp: "2025-06-15T12:00:00Z",
	}
	jsonData, _ := json.Marshal(tx)

	req := httptest.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		http.HandlerFunc(handlers.CreateTransactionHandler).ServeHTTP(rr, req)
		if rr.Code != http.StatusAccepted && rr.Code != http.StatusInternalServerError {
			b.Fatalf("Unexpected status code: %d", rr.Code)
		}
	}
}

func setupMockDB() *sql.DB {
	// Replace with mock DB connection or test DB
	// In production, use a mock SQL driver or test DB setup
	return nil
}
