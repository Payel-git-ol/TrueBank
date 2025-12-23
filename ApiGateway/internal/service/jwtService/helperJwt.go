package jwtService

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func getJwtKey() ([]byte, error) {
	jwtKey := []byte(os.Getenv("JWT_TOKEN"))

	if len(jwtKey) == 0 {
		return nil, errors.New("jwt token is empty")
	}

	return jwtKey, nil
}

func generateToken(name string, email string, jwtKey []byte) (string, error) {
	claims := jwt.MapClaims{
		"username": name,
		"email":    email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
