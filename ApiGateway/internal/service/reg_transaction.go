package service

import (
	"ApiGateway/internal/fetcher/kafka/producer"
	"ApiGateway/pkg/model"
)

func TransactionReg(data model.RegTransaction) error {
	regTransaction := model.RegTransaction{
		Name:                             data.Name,
		Description:                      data.Description,
		Company:                          data.Company,
		Documents:                        data.Documents,
		LinkToIndividualEntrepreneurship: data.LinkToIndividualEntrepreneurship,
	}

	err := producer.SendTransactionReg("server-reg", regTransaction)
	if err != nil {
		return err
	}

	return nil
}
