// finbank/fraud-service/main.go
package main

import (
	"log"

	"fraud-service/kafka"
	"fraud-service/rules"
	"fraud-service/alerts"
)

func main() {
	rules.LoadRules()
	alerts.InitAlertSystem()
	kafka.ConsumeTransactions()
}
