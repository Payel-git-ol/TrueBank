package message

import (
	"TrueBankUserService/pkg/cache"
	"TrueBankUserService/pkg/database"
	"TrueBankUserService/pkg/models"
	"TrueBankUserService/pkg/models/respons"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func ProcessMessageResultTransaction(data []byte) (respons.ResultTransaction, error) {
	fmt.Println("Consumer started")

	var res respons.ResultTransaction
	if err := json.Unmarshal(data, &res); err != nil {
		log.Printf("error unmarshalling: %v", err)
		return res, err
	}

	return res, nil
}

func ProcessMessageAuthCardNumber(data []byte) (models.AuthCardNumber, error) {
	fmt.Println("Consumer started")

	var res models.AuthCardNumber
	if err := json.Unmarshal(data, &res); err != nil {
		log.Println(err)
		return res, err
	}

	var user models.User
	if err := database.Db.Where("username = ?", res.Username).First(&user).Error; err != nil {
		log.Println(err)
		return res, err
	}

	user.CardNumber = strconv.Itoa(res.CardNumber)

	if err := database.Db.Model(&user).
		Updates(map[string]interface{}{"card_number": user.CardNumber}).Error; err != nil {
		log.Println("update error:", err)
		return res, err
	}

	fmt.Println("Card number updated for user:", user.Username)
	return res, nil
}

func ProcessMessageResultRemittance(data []byte) (respons.ResultRemittance, error) {
	fmt.Println("Consumer started")
	var res respons.ResultRemittance
	if err := json.Unmarshal(data, &res); err != nil {
		log.Println(err)
	}

	return res, nil
}

func ProcessMessageResultReplenishment(data []byte) (respons.Replenishment, error) {
	fmt.Println("Consumer started")

	var res respons.Replenishment
	if err := json.Unmarshal(data, &res); err != nil {
		log.Println(err)
	}

	if err := cache.ReplenishBalance(res.CardNumber, res.Sum); err != nil {
		log.Println("error replenishing balance:", err)
	}

	return res, nil
}
