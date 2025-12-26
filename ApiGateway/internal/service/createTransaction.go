package service

import (
	"ApiGateway/internal/kafkaService/producer"
	"ApiGateway/pkg/models"
)

func CreateTransaction(data models.TransactionRequest, nameTransaction string) error {
	transaction := models.Transaction{
		Username:        data.Username,
		NameTransaction: nameTransaction,
		Sum:             data.Sum,
		NumberCard:      data.NumberCard,
	}

	err := producer.SendTransaction("create-transaction", transaction)
	if err != nil {
		return err
	}

	return nil
}
