package consumer

import (
	"TrueBankUserService/internal/core/service"
	"TrueBankUserService/internal/core/service/message"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func GetAuthCardNumber(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "auth-card-number",
		GroupID: "get-auth-card-number",
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Message at offset %d: %s\n", len(m.Key), string(m.Value))
		msg, err := message.ProcessMessageAuthCardNumber(m.Value)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(msg)

		service.AuthCardNumberInCache(msg.Username, msg.CardNumber)
	}
}
