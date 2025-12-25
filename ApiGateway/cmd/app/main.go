package main

import (
	"ApiGateway/internal/grpc/auth"
	"ApiGateway/internal/grpc/authpb"
	"ApiGateway/internal/kafka/producer"
	"ApiGateway/internal/kafka/topic-generate/topic_init"
	"ApiGateway/internal/service/jwtService"
	"ApiGateway/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	r := gin.Default()
	var user models.User

	r.POST("/register", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := jwtService.UserServiceRegister(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		topic_init.InitTopic()
		if err := producer.SendMessageInRegistretion("register", user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		client := auth.NewAuthClient()
		req := &authpb.UserRequest{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			Balance:  user.Balance,
		}
		resp, err := auth.CallRegistration(client, req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  resp.Status,
			"message": resp.Message,
			"token":   token,
		})
	})

	r.POST("/auth", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		topic_init.InitTopic()
		if err := producer.SendMessageAuth("auth", user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		client := auth.NewAuthClient()
		req := &authpb.UserRequest{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			Balance:  user.Balance,
		}

		resp, err := auth.CallAuthentication(client, req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"status":  resp.Status,
			"message": resp.Message,
		})
	})

	r.GET("/profile/:id}", func(c *gin.Context) {
		//idParam := c.Param("id")

	})

	r.Run(":8080")
}
