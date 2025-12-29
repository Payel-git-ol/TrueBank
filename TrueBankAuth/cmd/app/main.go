package main

import (
	"TrueBankAuth/internal/fetcher/grpc/authpb"
	"TrueBankAuth/internal/fetcher/grpc/grpcinterceptor"
	"TrueBankAuth/internal/fetcher/grpc/server"
	consumer2 "TrueBankAuth/internal/fetcher/kafka/consumer"
	"TrueBankAuth/metrics"
	"TrueBankAuth/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

func main() {
	database.InitDb()
	var wg sync.WaitGroup
	wg.Add(1)
	go consumer2.GetMessageReg(&wg)
	go consumer2.GetMessageAuth(&wg)
	r := gin.Default()

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer(
			grpc.UnaryInterceptor(grpcinterceptor.MetricsInterceptor()),
		)
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

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	metrics.Init()

	r.Run(":7070")
}
