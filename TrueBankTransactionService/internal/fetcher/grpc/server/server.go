package server

import (
	transactionpb2 "TrueBankTransactionService/internal/fetcher/grpc/transactionpb"
	"TrueBankTransactionService/pkg/database"
	"TrueBankTransactionService/pkg/models"
	"context"
)

type TransactionServer struct {
	transactionpb2.UnimplementedTransactionServiceServer
}

func (s *TransactionServer) GetAllTransactions(ctx context.Context, req *transactionpb2.GetAllTransactionsRequest) (*transactionpb2.GetAllTransactionsResponse, error) {
	var txs []models.ListTransaction
	if err := database.Db.Find(&txs).Error; err != nil {
		return nil, err
	}

	resp := &transactionpb2.GetAllTransactionsResponse{Transactions: make([]*transactionpb2.TransactionItem, 0, len(txs))}
	for _, t := range txs {
		resp.Transactions = append(resp.Transactions, &transactionpb2.TransactionItem{
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
