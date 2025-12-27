package service

import (
	"ApiGateway/internal/kafkaService/producer/producer_transaction"
	"ApiGateway/pkg/models/transaction/reg"
)

func TransactionReg(data reg.RegTransaction) error {
	regTransaction := reg.RegTransaction{
		Name:                             data.Name,
		Description:                      data.Description,
		Company:                          data.Company,
		Documents:                        data.Documents,
		LinkToIndividualEntrepreneurship: data.LinkToIndividualEntrepreneurship,
	}

	err := producer_transaction.SendTransactionReg("server-reg", regTransaction)
	if err != nil {
		return err
	}

	return nil
}
