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

func GetResultRemittance(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
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

		metrics.KafkaMessagesOut.Inc()

		if err := service.UpdateUserInCacheRemittance(resultMessage.Username, resultMessage.SenderСardNumber, resultMessage.GetterCardNumber, resultMessage.Sum); err != nil {
			log.Println(err)
		}
	}
}
