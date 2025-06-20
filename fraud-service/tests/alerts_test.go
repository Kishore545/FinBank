package tests

import (
	"fraud/alerts"
	"testing"
)

func TestTriggerAlert(t *testing.T) {
	// This is a dummy test if TriggerAlert only prints/logs
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("TriggerAlert panicked: %v", r)
		}
	}()

	alertMessage := "Fraud detected in transaction ID 123"
	alerts.TriggerAlert(alertMessage)
}
