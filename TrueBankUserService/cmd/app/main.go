package main

import (
	"TrueBankUserService/internal/core/service"
	"TrueBankUserService/internal/fetcher/grpc/server"
	"TrueBankUserService/internal/fetcher/grpc/userservicepb"
	consumer2 "TrueBankUserService/internal/fetcher/kafka/consumer"
	"TrueBankUserService/pkg/cache"
	"TrueBankUserService/pkg/database"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

func main() {
	cache.InitRedis()
	database.InitDb()

	var wg sync.WaitGroup
	wg.Add(1)
	go consumer2.GetResultTransaction(&wg)
	go consumer2.GetAuthCardNumber(&wg)
	go consumer2.GetResultRemittance(&wg)
	go consumer2.GetMessageReplenishment(&wg)

	go func() {
		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		userservicepb.RegisterUserServiceServer(grpcServer, &server.UserServer{})

		log.Println("UserService started on :50052")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.GET("/search/profile/user/:username", func(c *gin.Context) {
		username := c.Param("username")

		cash, err := service.GetUserInCache(username)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "user not found in cache",
			})
			return
		}

		c.JSON(200, gin.H{
			"User": cash,
		})
	})

	r.Run(":5050")
}
