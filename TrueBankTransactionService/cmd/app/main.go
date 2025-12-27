package main

import (
	"TrueBankTransactionService/internal/kafkaService/consumer"
	"TrueBankTransactionService/pkg/database"
	"github.com/gin-gonic/gin"
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.Run(":6060")
}
