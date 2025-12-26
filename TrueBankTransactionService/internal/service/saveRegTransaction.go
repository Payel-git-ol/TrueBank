package service

import (
	"TrueBankTransactionService/pkg/database"
	"TrueBankTransactionService/pkg/models/dbModels"
	"TrueBankTransactionService/pkg/models/requestModels"
)

func SaveRegTransaction(data requestModels.RegTransaction) {
	newTransaction := dbModels.ListTransaction{
		Name:                             data.Name,
		Description:                      data.Description,
		Company:                          data.Company,
		Documents:                        data.Documents,
		LinkToIndividualEntrepreneurship: data.LinkToIndividualEntrepreneurship,
	}

	database.Db.Create(&newTransaction)
}
