package cache

import (
	"TrueBankUserService/pkg/models"
	"encoding/json"
	"time"
)

func SaveUser(user models.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return rdb.Set(ctx, user.Username, data, 30*24*time.Hour).Err()
}

func GetUser(username string) (*models.User, error) {
	val, err := rdb.Get(ctx, username).Result()
	if err != nil {
		return nil, err
	}
	var user models.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}
	return &user, nil
}
