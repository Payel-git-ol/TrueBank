package consumer

import (
	"TrueBankAuth/internal/core/service"
	"TrueBankAuth/internal/core/service/message"
	"TrueBankAuth/metrics"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func GetMessageAuth(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "server",
		GroupID: "auth-user-consumer",
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		result, err := message.ProcessMessage(m.Value)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)

		metrics.KafkaMessagesOut.Inc()

		service.AuthService(result)
	}
}
