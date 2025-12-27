package message

import (
	"TrueBankUserService/pkg/models"
	"encoding/json"
	"fmt"
	"log"
)

func ProcessMessageResultTransaction(data []byte) (models.ResultTransaction, error) {
	fmt.Println("Consumer started")

	var res models.ResultTransaction
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
	}

	return res, nil
}

func ProcessMessageResultRemittance(data []byte) (models.ResultRemittance, error) {
	fmt.Println("Consumer started")
	var res models.ResultRemittance
	if err := json.Unmarshal(data, &res); err != nil {
		log.Println(err)
	}

	return res, nil
}
