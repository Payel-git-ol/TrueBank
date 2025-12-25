package message

import (
	"TrueBankAuth/pkg/models"
	"encoding/json"
	"fmt"
	"log"
)

func ProcessMessage(data []byte) (models.RequestUser, error) {
	fmt.Println("Consumer started")

	var req models.RequestUser
	if err := json.Unmarshal(data, &req); err != nil {
		log.Fatal(err)
	}

	return req, nil
}
