package handlers

import "transaction/models"

type DB interface {
	InsertTransaction(tx models.Transaction) error
}
