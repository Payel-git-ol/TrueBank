package service

import (
	"TrueBankAuth/pkg/database"
	"TrueBankAuth/pkg/models"
	"log"
)

func RegService(req models.RequestUser) {
	newUser := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Balance:  req.Balance,
		Role:     "user",
	}

	var existing models.User
	result := database.Db.Where("username = ? AND email = ?", req.Username, req.Email).First(&existing)

	if result.Error == nil {
		log.Println("User already exists:", existing.Username)
		return
	}

	if err := database.Db.Create(&newUser).Error; err != nil {
		log.Println("Error creating user:", err)
	}
}
