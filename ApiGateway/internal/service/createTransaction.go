package service

import (
	"ApiGateway/internal/kafkaService/producer"
	"ApiGateway/pkg/models"
	"log"
)

func CreateTransaction(data models.TransactionRequest, nameTransaction string) {
	transaction := models.Transaction{
		Username:        data.Username,
		NameTransaction: nameTransaction,
		Sum:             data.Sum,
		NumberCard:      data.NumberCard,
	}

	err := producer.SendTransaction("create-transaction", transaction)
	if err != nil {
		log.Println(err)
	}
}
