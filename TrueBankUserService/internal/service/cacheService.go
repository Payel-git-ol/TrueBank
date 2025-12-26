package service

import (
	"TrueBankUserService/pkg/cache"
	"TrueBankUserService/pkg/models"
	"strconv"
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

func AuthCardNumber(username string, cardNumber string) error {
	cardNumberInt, _ := strconv.Atoi(cardNumber)

	err := cache.AuthCardNumber(username, cardNumberInt)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserInCash(CardNumber string, subtractAmount float64) error {
	err := cache.UpdateUser(CardNumber, subtractAmount)
	if err != nil {
		return err
	}

	return nil
}

func TestAddBalance(username string, amount float64) error {
	err := cache.AddBalance(username, amount)
	if err != nil {
		return err
	}

	return nil
}
