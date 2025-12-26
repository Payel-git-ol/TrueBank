package service

import (
	"TrueBankTransactionService/internal/kafkaService/producer"
	"TrueBankTransactionService/pkg/database"
	"TrueBankTransactionService/pkg/models/dbModels"
	"strconv"
)

func CreateTransaction(newTransaction dbModels.HistoryTransaction) error {
	database.Db.Create(&newTransaction)

	sumFloat, err := strconv.ParseFloat(newTransaction.Sum, 64)
	if err != nil {
		return err
	}

	err = producer.SendMessageTransaction("result-transaction", sumFloat, newTransaction.NumberCard, newTransaction.Username)
	if err != nil {
		return err
	}

	return nil
}
