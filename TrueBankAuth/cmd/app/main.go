package main

import (
	"TrueBankAuth/internal/grpc/authpb"
	"TrueBankAuth/internal/grpc/server"
	"TrueBankAuth/internal/kafka/consumer"
	"TrueBankAuth/pkg/database"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

func main() {
	database.InitDb()
	var wg sync.WaitGroup
	wg.Add(1)
	go consumer.GetMessageReg(&wg)
	go consumer.GetMessageAuth(&wg)
	r := gin.Default()

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		authpb.RegisterAuthServiceServer(grpcServer, &server.AuthServer{})
		log.Println("gRPC server started on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.Run(":7070")
}
