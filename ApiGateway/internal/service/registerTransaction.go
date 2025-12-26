package service

import (
	"ApiGateway/internal/kafkaService/producer"
	"ApiGateway/pkg/models"
)

func TransactionReg(data models.RegTransaction) error {
	regTransaction := models.RegTransaction{
		Name:                             data.Name,
		Description:                      data.Description,
		Company:                          data.Company,
		Documents:                        data.Documents,
		LinkToIndividualEntrepreneurship: data.LinkToIndividualEntrepreneurship,
	}

	err := producer.SendTransactionReg("transaction-reg", regTransaction)
	if err != nil {
		return err
	}

	return nil
}
