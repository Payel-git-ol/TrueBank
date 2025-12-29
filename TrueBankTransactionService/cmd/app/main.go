package main

import (
	"TrueBankTransactionService/internal/fetcher/grpc/server"
	"TrueBankTransactionService/internal/fetcher/grpc/transactionpb"
	consumer2 "TrueBankTransactionService/internal/fetcher/kafka/consumer"
	"TrueBankTransactionService/pkg/database"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

func main() {

	database.InitDb()
	r := gin.Default()

	var wg sync.WaitGroup
	wg.Add(1)
	go consumer2.GetTransaction(&wg)
	go consumer2.GetRegTransaction(&wg)
	go consumer2.GetMessageRemittance(&wg)

	go func() {
		lis, err := net.Listen("tcp", ":50053")

		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		transactionpb.RegisterTransactionServiceServer(grpcServer, &server.TransactionServer{})
		reflection.Register(grpcServer)
		log.Println("TransactionService gRPC server started on :50053")

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.Run(":6060")
}
