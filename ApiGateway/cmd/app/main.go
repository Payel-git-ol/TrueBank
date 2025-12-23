package main

import (
	"ApiGateway/internal/service/jwtService"
	"ApiGateway/pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var user models.User

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.POST("/new/register", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		jwtService.UserServiceRegister(user)
	})

	r.Run(":8080")
}
