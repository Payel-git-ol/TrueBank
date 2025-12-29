package message

import (
	"TrueBankTransactionService/pkg/models"
	"TrueBankTransactionService/pkg/models/requests"
	"encoding/json"
	"fmt"
	"log"
)

func ProcessMessageNewTransaction(data []byte) (requests.TransactionRequest, error) {
	fmt.Println("Consumer started")

	var req requests.TransactionRequest
	if err := json.Unmarshal(data, &req); err != nil {
		log.Fatal(err)
	}

	return req, nil
}

func ProcessMessageRegTransaction(data []byte) (models.RegTransaction, error) {
	fmt.Println("Consumer started")

	var req models.RegTransaction
	if err := json.Unmarshal(data, &req); err != nil {
		log.Fatal(err)
	}

	return req, nil
}

func ProcessMessageRemittance(data []byte) (requests.RemittanceRequest, error) {
	fmt.Println("Consumer started")

	var req requests.RemittanceRequest
	if err := json.Unmarshal(data, &req); err != nil {
		log.Fatal(err)
	}

	return req, nil
}
