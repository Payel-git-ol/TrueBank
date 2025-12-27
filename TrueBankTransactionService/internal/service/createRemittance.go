package service

import (
	"TrueBankTransactionService/internal/kafkaService/producer"
	"TrueBankTransactionService/pkg/database"
	"TrueBankTransactionService/pkg/models/dbModels"
	"TrueBankTransactionService/pkg/models/requestModels"
)

func CreateRemittance(data dbModels.RemittanceHistory) {
	database.Db.Create(&data)

	result := requestModels.ResultRemittance{
		Username:         data.Username,
		SenderСardNumber: data.SenderСardNumber,
		GetterCardNumber: data.GetterCardNumber,
		Sum:              data.Sum,
	}

	producer.SendMessageRemittance("result-remittance", result)
}
