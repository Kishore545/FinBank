// finbank/fraud-service/tests/alerts_test.go
package tests

import (
	"testing"
	"fraud-service/alerts"
)

func TestTriggerAlert(t *testing.T) {
	alertMessage := "Fraud detected in transaction ID 123"
	output := alerts.TriggerAlert(alertMessage)
	
	if output != "Alert sent: "+alertMessage {
		t.Errorf("Expected alert message not received, got %s", output)
	}
}
