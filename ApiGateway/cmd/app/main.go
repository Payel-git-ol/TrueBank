package main

import (
	"ApiGateway/internal/grpc/auth"
	"ApiGateway/internal/grpc/authpb"
	"ApiGateway/internal/kafkaService/producer"
	"ApiGateway/internal/kafkaService/topic-generate/topic_init"
	"ApiGateway/internal/service"
	"ApiGateway/internal/service/jwtService"
	"ApiGateway/pkg/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
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
		if err := producer.SendMessageAuth("server", user); err != nil {
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

	r.GET("/profile/:username", func(c *gin.Context) {
		username := c.Param("username")

		resp, err := http.Get("http://localhost:5050/search/profile/user/" + username)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		var userResp models.UserResponse
		if err := json.Unmarshal(body, &userResp); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"status": resp.Status,
			"user":   userResp.User,
		})
	})

	r.POST("/payment/service/:name", func(c *gin.Context) {
		name := c.Param("name")

		var transaction models.TransactionRequest

		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		service.CreateTransaction(transaction, name)
	})

	r.POST("/payment/reg", func(c *gin.Context) {
		var regTransaction models.RegTransaction

		if err := c.ShouldBindJSON(&regTransaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		service.TransactionReg(regTransaction)
	})

	r.Run(":8080")
}
