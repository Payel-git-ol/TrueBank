package server

import (
	"TrueBankTransactionService/internal/grpc/transactionpb"
	"TrueBankTransactionService/pkg/database"
	"TrueBankTransactionService/pkg/models/dbModels"
	"context"
)

type TransactionServer struct {
	transactionpb.UnimplementedTransactionServiceServer
}

func (s *TransactionServer) GetAllTransactions(ctx context.Context, req *transactionpb.GetAllTransactionsRequest) (*transactionpb.GetAllTransactionsResponse, error) {
	var txs []dbModels.ListTransaction
	if err := database.Db.Find(&txs).Error; err != nil {
		return nil, err
	}

	resp := &transactionpb.GetAllTransactionsResponse{Transactions: make([]*transactionpb.TransactionItem, 0, len(txs))}
	for _, t := range txs {
		resp.Transactions = append(resp.Transactions, &transactionpb.TransactionItem{
			Id:                               t.Id,
			Name:                             t.Name,
			Description:                      t.Description,
			Company:                          t.Company,
			Documents:                        t.Documents,
			LinkToIndividualEntrepreneurship: t.LinkToIndividualEntrepreneurship,
		})
	}
	return resp, nil
}
