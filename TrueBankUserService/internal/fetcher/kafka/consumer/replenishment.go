package consumer

import (
	"TrueBankUserService/internal/core/service/message"
	"TrueBankUserService/metrics"
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func GetMessageReplenishment(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "replenishment",
		GroupID: "group-replenishment",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		metrics.KafkaMessagesOut.Inc()

		metrics.KafkaMessagesOut.Inc()

		message.ProcessMessageResultReplenishment(msg.Value)
	}
}
