package consumer

import (
	"TrueBankTransactionService/internal/kafkaService/message"
	"TrueBankTransactionService/internal/service"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func GetRegTransaction(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "transaction-reg",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		processMessage, err := message.ProcessMessageRegTransaction(msg.Value)
		if err != nil {
			log.Fatal(err)
		}

		service.SaveRegTransaction(processMessage)

		fmt.Println(processMessage)
	}
}
