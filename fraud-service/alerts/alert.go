// finbank/fraud-service/alerts/alert.go
package alerts

import (
	"fmt"
	"log"
)

type Transaction struct {
	ID        int
	UserID    int
	Amount    float64
	Timestamp string
}

func InitAlertSystem() {
	log.Println("Alert system initialized")
}

func SendAlert(tx Transaction) {
	msg := fmt.Sprintf("ALERT: Fraudulent transaction detected! ID: %d, User: %d, Amount: %.2f", tx.ID, tx.UserID, tx.Amount)
	log.Println(msg)
}
