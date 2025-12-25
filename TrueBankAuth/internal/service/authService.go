package service

import (
	"TrueBankAuth/pkg/database"
	"TrueBankAuth/pkg/models"
)

func AuthService(req models.RequestUser) (string, error) {
	authUser := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Balance:  req.Balance,
		Role:     "user",
	}

	result := database.Db.Where("username = ? AND email = ?", req.Username, req.Email).First(&authUser)
	if result.Error != nil {
		return "No user found", result.Error
	}

	return "User succes", nil
}
