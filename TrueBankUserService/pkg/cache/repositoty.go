package cache

import (
	"TrueBankUserService/pkg/models"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func SaveUser(user models.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return rdb.Set(ctx, "user:"+user.Username, data, 30*24*time.Hour).Err()
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

func AuthCardNumber(username string, cardNumber int) error {
	ctx := context.Background()
	key := "user:" + username

	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return errors.New("user not found")
	} else if err != nil {
		return err
	}

	var user models.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return err
	}

	user.CardNumber = strconv.Itoa(cardNumber)

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, key, data, 24*time.Hour).Err()
}

func UpdateUser(username string, subtractAmount float64) error {
	ctx := context.Background()
	key := "user:" + username

	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return errors.New("user not found")
	} else if err != nil {
		return err
	}

	var user models.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return err
	}

	user.Balance -= subtractAmount
	if user.Balance < 0 {
		return errors.New("user balance is negative")
	}

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, key, data, 24*time.Hour).Err()
}

func AddBalance(username string, amount float64) error {
	ctx := context.Background()
	key := "user:" + username

	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return errors.New("user not found")
	} else if err != nil {
		return err
	}

	var user models.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return err
	}

	user.Balance += amount

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, key, data, 24*time.Hour).Err()
}
