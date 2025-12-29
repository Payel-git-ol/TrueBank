package jwt

import (
	"ApiGateway/pkg/model"
)

func UserServiceRegister(user model.User) (string, error) {
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
