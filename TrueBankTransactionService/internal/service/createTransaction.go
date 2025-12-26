package service

import (
	"TrueBankTransactionService/pkg/database"
	"TrueBankTransactionService/pkg/models/dbModels"
	"TrueBankTransactionService/pkg/models/requestModels"
	"time"
)

func CreateTransaction(data requestModels.TransactionRequest) {
	newTransaction := dbModels.HistoryTransaction{
		Username:        data.Username,
		NameTransaction: data.NameTransaction,
		Sum:             data.Sum,
		NumberCard:      data.NumberCard,
		DateCreated:     time.Now(),
	}

	database.Db.Create(&newTransaction)
}
