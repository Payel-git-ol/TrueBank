package service

import (
	"TrueBankTransactionService/pkg/database"
	"TrueBankTransactionService/pkg/models/dbModels"
)

func SaveRegTransaction(data dbModels.ListTransaction) {
	database.Db.Create(data)
}
