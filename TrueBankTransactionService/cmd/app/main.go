package main

import (
	"TrueBankTransactionService/internal/grpc/server"
	"TrueBankTransactionService/internal/grpc/transactionpb"
	"TrueBankTransactionService/internal/kafkaService/consumer"
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
	go consumer.GetTransaction(&wg)
	go consumer.GetRegTransaction(&wg)
	go consumer.GetMessageRemittance(&wg)

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
