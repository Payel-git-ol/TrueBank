package main

import (
	"ApiGateway/internal/fetcher/grpc/auth"
	"ApiGateway/internal/fetcher/grpc/authpb"
	"ApiGateway/internal/fetcher/grpc/client"
	producer2 "ApiGateway/internal/fetcher/kafka/producer"
	"ApiGateway/internal/service"
	"ApiGateway/internal/service/jwt"
	"ApiGateway/pkg/model"
	"ApiGateway/pkg/model/requests"
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
	var user model.User

	r.POST("/register", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := jwt.UserServiceRegister(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := producer2.SendMessageInRegistretion("register", user); err != nil {
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

		if err := producer2.SendMessageAuth("server", user); err != nil {
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

	r.POST("/auth/cardNumber", func(c *gin.Context) {
		var authCardNumber model.AuthCardNumber
		if err := c.ShouldBindJSON(&authCardNumber); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		err := producer2.SendMessageAuthCardNumber("auth-card-number", authCardNumber)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		c.JSON(200, gin.H{
			"status":  200,
			"message": "success",
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

		var userResp model.UserResponse
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

		var transaction requests.TransactionRequest

		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err := service.CreateTransaction(transaction, name)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		c.JSON(200, gin.H{
			"status":  200,
			"message": "success",
		})
	})

	r.GET("/transactions", func(c *gin.Context) {
		grpcClient := client.NewTransactionClient()

		resp, err := client.CallGetAllTransactions(grpcClient)
		if err != nil {
			log.Println("GetAllTransactions error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transactions"})
			return
		}

		c.JSON(http.StatusOK, resp.Transactions)
	})

	r.POST("/payment/reg", func(c *gin.Context) {
		var regTransaction model.RegTransaction

		if err := c.ShouldBindJSON(&regTransaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err := service.TransactionReg(regTransaction)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		c.JSON(200, gin.H{
			"status":  200,
			"message": "success",
		})
	})

	r.POST("/remittance", func(c *gin.Context) {
		var remittanceTransaction model.RemittanceTransaction

		if err := c.ShouldBindJSON(&remittanceTransaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err := producer2.SendMessageRemittance("create-remittance", remittanceTransaction)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(200, gin.H{
			"status":  200,
			"message": "success",
		})
	})

	r.POST("/replenishment", func(c *gin.Context) {
		var replenishment requests.Replenishment
		if err := c.ShouldBindJSON(&replenishment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		producer2.SendMessageReplenishment("replenishment", replenishment)
	})

	r.Run(":8080")
}
