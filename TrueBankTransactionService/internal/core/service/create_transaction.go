package service

import (
	"TrueBankTransactionService/internal/fetcher/kafka/producer"
	"TrueBankTransactionService/pkg/database"
	"TrueBankTransactionService/pkg/models"
	"strconv"
)

func CreateTransaction(newTransaction models.HistoryTransaction) error {
	database.Db.Create(&newTransaction)

	sumFloat, err := strconv.ParseFloat(newTransaction.Sum, 64)
	if err != nil {
		return err
	}

	err = producer.SendMessageTransaction("result-server", sumFloat, newTransaction.NumberCard, newTransaction.Username)
	if err != nil {
		return err
	}

	return nil
}
