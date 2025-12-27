package service

import (
	"ApiGateway/internal/kafkaService/producer/producer_transaction"
	"ApiGateway/pkg/models/transaction"
	"ApiGateway/pkg/models/transaction/request"
)

func CreateTransaction(data request.TransactionRequest, nameTransaction string) error {
	transaction := transaction.Transaction{
		Username:        data.Username,
		NameTransaction: nameTransaction,
		Sum:             data.Sum,
		NumberCard:      data.NumberCard,
	}

	err := producer_transaction.SendTransaction("create-server", transaction)
	if err != nil {
		return err
	}

	return nil
}
