package service

import (
	"ApiGateway/internal/fetcher/kafka/producer"
	"ApiGateway/pkg/model"
	"ApiGateway/pkg/model/requests"
)

func CreateTransaction(data requests.TransactionRequest, nameTransaction string) error {
	transaction := model.Transaction{
		Username:        data.Username,
		NameTransaction: nameTransaction,
		Sum:             data.Sum,
		NumberCard:      data.NumberCard,
	}

	err := producer.SendTransaction("create-server", transaction)
	if err != nil {
		return err
	}

	return nil
}
