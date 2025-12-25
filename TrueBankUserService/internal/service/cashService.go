package service

import (
	"TrueBankUserService/pkg/cache"
	"TrueBankUserService/pkg/models"
)

func SaveUserInCash(user models.User) error {
	err := cache.SaveUser(user)
	if err != nil {
		return err
	}

	return nil
}

func GetUserInCash(key string) (*models.User, error) {
	get, err := cache.GetUser(key)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: get.Username,
		Email:    get.Email,
		Role:     get.Role,
	}

	return user, nil
}
