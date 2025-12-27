package consumer

import (
	"TrueBankUserService/internal/kafka/consumer/message"
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func GetMessageReplenishment(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "replenishment",
		GroupID: "group-replenishment",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		message.ProcessMessageResultReplenishment(msg.Value)
	}
}
