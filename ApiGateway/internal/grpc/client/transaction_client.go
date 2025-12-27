package client

import (
	"ApiGateway/internal/grpc/transactionpb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func NewTransactionClient() transactionpb.TransactionServiceClient {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return transactionpb.NewTransactionServiceClient(conn)
}

func CallGetAllTransactions(client transactionpb.TransactionServiceClient) (*transactionpb.GetAllTransactionsResponse, error) {
	return client.GetAllTransactions(context.Background(), &transactionpb.GetAllTransactionsRequest{})
}
