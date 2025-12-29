package consumer

import (
	"TrueBankUserService/internal/core/service"
	"TrueBankUserService/internal/core/service/message"
	"TrueBankUserService/metrics"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func GetResultTransaction(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "result-server",
		GroupID: "get-res-server",
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

		metrics.KafkaMessagesOut.Inc()

		if err := service.UpdateUserInCacheTransaction(resultMessage.Username, resultMessage.Sum); err != nil {
			log.Printf("error updating user: %v", err)
		}
	}
}
