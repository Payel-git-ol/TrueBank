package service

import (
	"TrueBankTransactionService/pkg/database"
	"TrueBankTransactionService/pkg/models"
)

func SaveRegTransaction(data models.ListTransaction) {
	database.Db.Create(data)
}
