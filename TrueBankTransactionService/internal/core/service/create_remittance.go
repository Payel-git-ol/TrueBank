package service

import (
	"TrueBankTransactionService/internal/fetcher/kafka/producer"
	"TrueBankTransactionService/pkg/database"
	"TrueBankTransactionService/pkg/models"
	"TrueBankTransactionService/pkg/models/respons"
)

func CreateRemittance(data models.RemittanceHistory) {
	database.Db.Create(&data)

	result := respons.ResultRemittance{
		Username:         data.Username,
		SenderСardNumber: data.SenderСardNumber,
		GetterCardNumber: data.GetterCardNumber,
		Sum:              data.Sum,
	}

	producer.SendMessageRemittance("result-remittance", result)
}
