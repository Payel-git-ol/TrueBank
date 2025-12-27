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

	pipe := rdb.TxPipeline()
	pipe.Set(ctx, key, data, 24*time.Hour)
	pipe.Set(ctx, "card:"+strconv.Itoa(cardNumber), user.Username, 24*time.Hour)
	_, err = pipe.Exec(ctx)
	return err

}

func UpdateUserTransaction(username string, subtractAmount float64) error {
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

func UpdateUserRemittance(senderUsername string, senderCardNumber string, getterCardNumber string, amount float64) error {
	ctx := context.Background()

	senderKey := "user:" + senderUsername
	senderVal, err := rdb.Get(ctx, senderKey).Result()

	if err == redis.Nil {
		return errors.New("sender not found")
	} else if err != nil {
		return err
	}

	var sender models.User
	if err := json.Unmarshal([]byte(senderVal), &sender); err != nil {
		return err
	}

	if sender.CardNumber != senderCardNumber {
		return errors.New("sender card mismatch")
	}

	getterUsername, err := rdb.Get(ctx, "card:"+getterCardNumber).Result()

	if err == redis.Nil {
		return errors.New("getter not found")
	} else if err != nil {
		return err
	}

	getterKey := "user:" + getterUsername
	getterVal, err := rdb.Get(ctx, getterKey).Result()

	if err == redis.Nil {
		return errors.New("getter user not found")
	} else if err != nil {
		return err
	}

	var getter models.User
	if err := json.Unmarshal([]byte(getterVal), &getter); err != nil {
		return err
	}

	if sender.Balance < amount {
		return errors.New("insufficient funds")
	}

	sender.Balance -= amount

	getter.Balance += amount

	senderData, _ := json.Marshal(sender)
	getterData, _ := json.Marshal(getter)

	pipe := rdb.TxPipeline()
	pipe.Set(ctx, senderKey, senderData, 24*time.Hour)
	pipe.Set(ctx, getterKey, getterData, 24*time.Hour)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
