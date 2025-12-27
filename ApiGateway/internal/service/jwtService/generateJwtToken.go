package jwtService

import (
	"ApiGateway/pkg/models/user"
)

func UserServiceRegister(user user.User) (string, error) {
	jwtKey, err := getJwtKey()
	if err != nil {
		return "Error: ", err
	}

	result, err := generateToken(user.Username, user.Email, jwtKey)
	if err != nil {
		return "Error: ", err
	}

	return "Success" + result, err
}
