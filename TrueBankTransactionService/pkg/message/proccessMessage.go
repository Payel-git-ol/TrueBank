package message

import (
	"TrueBankTransactionService/pkg/models/requestModels"
	"encoding/json"
	"fmt"
	"log"
)

func ProcessMessageNewTransaction(data []byte) (requestModels.TransactionRequest, error) {
	fmt.Println("Consumer started")

	var req requestModels.TransactionRequest
	if err := json.Unmarshal(data, &req); err != nil {
		log.Fatal(err)
	}

	return req, nil
}

func ProcessMessageRegTransaction(data []byte) (requestModels.RegTransaction, error) {
	fmt.Println("Consumer started")

	var req requestModels.RegTransaction
	if err := json.Unmarshal(data, &req); err != nil {
		log.Fatal(err)
	}

	return req, nil
}
