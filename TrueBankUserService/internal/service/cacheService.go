package service

import (
	"TrueBankUserService/pkg/cache"
	"TrueBankUserService/pkg/models"
)

func SaveUserInCache(user models.User) error {
	err := cache.SaveUser(user)
	if err != nil {
		return err
	}

	return nil
}

func GetUserInCache(key string) (*models.User, error) {
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

func AuthCardNumberInCache(username string, cardNumber int) error {
	err := cache.AuthCardNumber(username, cardNumber)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserInCacheTransaction(CardNumber string, subtractAmount float64) error {
	err := cache.UpdateUserTransaction(CardNumber, subtractAmount)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserInCacheRemittance(username string, SenderСardNumber string, GetterCardNumber string, amount float64) error {
	err := cache.UpdateUserRemittance(username, SenderСardNumber, GetterCardNumber, amount)
	if err != nil {
		return err
	}

	return nil
}
