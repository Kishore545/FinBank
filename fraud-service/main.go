package main

import (
	"fraud/kafka"
	"fraud/rules"
)

func main() {
	rules.LoadRules()
	kafka.ConsumeTransactions()
}
