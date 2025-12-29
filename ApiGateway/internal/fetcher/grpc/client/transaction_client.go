package client

import (
	transactionpb2 "ApiGateway/internal/fetcher/grpc/transactionpb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func NewTransactionClient() transactionpb2.TransactionServiceClient {
	conn, err := grpc.Dial("transaction-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return transactionpb2.NewTransactionServiceClient(conn)
}

func CallGetAllTransactions(client transactionpb2.TransactionServiceClient) (*transactionpb2.GetAllTransactionsResponse, error) {
	return client.GetAllTransactions(context.Background(), &transactionpb2.GetAllTransactionsRequest{})
}
