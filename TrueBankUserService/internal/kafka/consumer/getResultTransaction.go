package consumer

import (
	"TrueBankUserService/internal/kafka/consumer/message"
	"TrueBankUserService/internal/service"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func GetResultTransaction(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "result-transaction",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Message at offset %d: %s\n", msg.Offset, string(msg.Value))

		resultMessage, err := message.ProcessMessageResultTransaction(msg.Value)
		if err != nil {
			log.Fatal(err)
		}

		if err := service.UpdateUserInCash(resultMessage.Username, resultMessage.Sum); err != nil {
			log.Printf("error updating user: %v", err)
		}
	}
}
