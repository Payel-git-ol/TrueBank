package jwtService

import (
	"ApiGateway/pkg/models"
)

func UserServiceRegister(user models.User) (string, error) {
	jwtKey, err := getJwtKey()
	if err != nil {
		return "Error: ", err
	}

	generateToken(user.Username, user.Email, jwtKey)
	return "Success", err
}
