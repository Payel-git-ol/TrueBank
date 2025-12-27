package cache

import (
	"TrueBankUserService/pkg/database"
	"TrueBankUserService/pkg/models"
	"context"
	"encoding/json"
	"time"
)

func ReplenishBalance(cardNumber string, sum float64) error {
	ctx := context.Background()

	var user models.User
	if err := database.Db.Where("card_number = ?", cardNumber).First(&user).Error; err != nil {
		return err
	}

	user.Balance += sum
	if err := database.Db.Save(&user).Error; err != nil {
		return err
	}

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	pipe := rdb.TxPipeline()
	pipe.Set(ctx, "user:"+user.Username, data, time.Hour) // TTL = 1 час
	pipe.Set(ctx, "card:"+user.CardNumber, user.Username, time.Hour)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
