package consumer

import (
	"TrueBankTransactionService/internal/service"
	"TrueBankTransactionService/pkg/message"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func GetTransaction(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "create-transaction",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		messageResult, err := message.ProcessMessageNewTransaction(msg.Value)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(messageResult)

		service.CreateTransaction(messageResult)
	}
}
