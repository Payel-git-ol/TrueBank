package consumer

import (
	"TrueBankTransactionService/internal/kafkaService/message"
	"TrueBankTransactionService/internal/service"
	"TrueBankTransactionService/pkg/models/dbModels"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
	"time"
)

func GetMessageRemittance(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "create-remittance",
		GroupID: "get-remittance",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
		}

		messageResult, err := message.ProcessMessageRemittance(msg.Value)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(messageResult)

		newRemittance := dbModels.RemittanceHistory{
			Username:         messageResult.Username,
			SenderСardNumber: messageResult.SenderСardNumber,
			GetterCardNumber: messageResult.GetterCardNumber,
			Sum:              messageResult.Sum,
			DataCreate:       time.Now(),
		}

		service.CreateRemittance(newRemittance)
	}
}
