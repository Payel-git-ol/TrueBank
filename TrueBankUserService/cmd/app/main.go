package main

import (
	"TrueBankUserService/internal/grpc/server"
	"TrueBankUserService/internal/grpc/userservicepb"
	"TrueBankUserService/internal/service"
	"TrueBankUserService/pkg/cache"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cache.InitRedis()

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

		cash, err := service.GetUserInCash(username)
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
