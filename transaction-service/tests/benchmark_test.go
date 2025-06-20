package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"transaction/handlers"
	"transaction/models"
)

// mockDB implements handlers.DB
type mockDB struct{}

func (m *mockDB) InsertTransaction(tx models.Transaction) error {
	// Simulate DB insert logic (no-op)
	return nil
}

func BenchmarkCreateTransactionHandler(b *testing.B) {
	handlers.Init(&mockDB{})

	tx := models.Transaction{
		ID:        1,
		UserID:    101,
		Amount:    123.45,
		Timestamp: "2025-06-15T12:00:00Z",
	}
	jsonData, _ := json.Marshal(tx)

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		http.HandlerFunc(handlers.CreateTransactionHandler).ServeHTTP(rr, req)

		if rr.Code != http.StatusAccepted && rr.Code != http.StatusInternalServerError {
			b.Fatalf("Unexpected status code: %d", rr.Code)
		}
	}
}
