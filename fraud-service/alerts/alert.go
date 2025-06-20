package alerts

import (
	"fmt"
	"fraud/models"
)

// TriggerAlert is used in alerts_test.go
func TriggerAlert(message string) string {
	alertMsg := "Alert sent: " + message
	fmt.Println(alertMsg)
	return alertMsg
}

// SendAlert is used in consumer.go
func SendAlert(tx models.Transaction) {
	msg := fmt.Sprintf("Fraud detected for user %d on transaction ID %d", tx.UserID, tx.ID)
	TriggerAlert(msg)
}
