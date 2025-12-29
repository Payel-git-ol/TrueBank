package service

import (
	"TrueBankUserService/pkg/cache"
	"TrueBankUserService/pkg/database"
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
	return cache.rdb.Set(cache.ctx, "user:"+user.Username, data, 30*24*time.Hour).Err()
}

func GetUser(username string) (*models.User, error) {
	ctx := context.Background()
	key := "user:" + username

	val, err := cache.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		// если нет в кэше — берём из БД
		var user models.User
		if err := database.Db.Where("username = ?", username).First(&user).Error; err != nil {
			return nil, err
		}

		data, _ := json.Marshal(user)
		_ = cache.rdb.Set(ctx, key, data, time.Hour).Err()

		return &user, nil
	} else if err != nil {
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

	val, err := cache.rdb.Get(ctx, key).Result()
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

	pipe := cache.rdb.TxPipeline()
	pipe.Set(ctx, key, data, 24*time.Hour)
	pipe.Set(ctx, "card:"+strconv.Itoa(cardNumber), user.Username, 24*time.Hour)
	_, err = pipe.Exec(ctx)
	return err

}

func UpdateUserTransaction(username string, sum float64) error {
	ctx := context.Background()
	key := "user:" + username

	val, err := cache.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return errors.New("user not found in cache")
	} else if err != nil {
		return err
	}

	var user models.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return err
	}

	user.Balance -= sum
	if user.Balance < 0 {
		return errors.New("insufficient funds")
	}

	data, _ := json.Marshal(user)
	if err := cache.rdb.Set(ctx, key, data, time.Hour).Err(); err != nil {
		return err
	}

	if err := database.Db.Model(&models.User{}).
		Where("username = ?", username).
		Update("balance", user.Balance).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUserRemittance(senderUsername, senderCardNumber, getterCardNumber string, amount float64) error {
	ctx := context.Background()

	senderVal, err := cache.rdb.Get(ctx, "user:"+senderUsername).Result()
	if err != nil {
		return err
	}
	var sender models.User
	_ = json.Unmarshal([]byte(senderVal), &sender)

	if sender.CardNumber != senderCardNumber {
		return errors.New("sender card mismatch")
	}

	getterUsername, err := cache.rdb.Get(ctx, "card:"+getterCardNumber).Result()
	if err != nil {
		return err
	}
	getterVal, _ := cache.rdb.Get(ctx, "user:"+getterUsername).Result()
	var getter models.User
	_ = json.Unmarshal([]byte(getterVal), &getter)

	if sender.Balance < amount {
		return errors.New("insufficient funds")
	}

	sender.Balance -= amount
	getter.Balance += amount

	senderData, _ := json.Marshal(sender)
	getterData, _ := json.Marshal(getter)

	pipe := cache.rdb.TxPipeline()
	pipe.Set(ctx, "user:"+senderUsername, senderData, time.Hour)
	pipe.Set(ctx, "user:"+getterUsername, getterData, time.Hour)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}

	if err := database.Db.Model(&models.User{}).
		Where("username = ?", senderUsername).
		Update("balance", sender.Balance).Error; err != nil {
		return err
	}
	if err := database.Db.Model(&models.User{}).
		Where("username = ?", getterUsername).
		Update("balance", getter.Balance).Error; err != nil {
		return err
	}

	return nil
}
