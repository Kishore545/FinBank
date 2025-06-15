// finbank/transaction-service/tests/transaction_test.go
package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"transaction-service/handlers"
)

func TestCreateTransactionHandler(t *testing.T) {
	reqBody := []byte(`{"user_id":1,"amount":99.99,"timestamp":"2024-06-15T10:00:00Z"}`)
	req := httptest.NewRequest("POST", "/transaction", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.CreateTransactionHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("Expected status 202 Accepted, got %v", rr.Code)
	}
}
