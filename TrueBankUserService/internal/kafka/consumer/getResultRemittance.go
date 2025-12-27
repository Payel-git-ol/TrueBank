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

func GetResultRemittance(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "result-remittance",
		GroupID: "get-res-remittance",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("Message at offset %d: %s\n", msg.Offset, string(msg.Value))

		resultMessage, err := message.ProcessMessageResultRemittance(msg.Value)
		if err != nil {
			log.Println(err)
		}

		if err := service.UpdateUserInCacheRemittance(resultMessage.Username, resultMessage.SenderСardNumber, resultMessage.GetterCardNumber, resultMessage.Sum); err != nil {
			log.Println(err)
		}
	}
}
