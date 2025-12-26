package service

import (
	"ApiGateway/internal/kafkaService/producer"
	"ApiGateway/pkg/models"
)

func TransactionReg(data models.RegTransaction) {
	regTransaction := models.RegTransaction{
		Name:                             data.Name,
		Description:                      data.Description,
		Company:                          data.Company,
		Documents:                        data.Documents,
		LinkToIndividualEntrepreneurship: data.LinkToIndividualEntrepreneurship,
	}

	producer.SendTransactionReg("transaction-reg", regTransaction)
}
