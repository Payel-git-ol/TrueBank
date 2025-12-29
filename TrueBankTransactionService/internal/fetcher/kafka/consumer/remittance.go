package consumer

import (
	"TrueBankTransactionService/internal/core/service"
	"TrueBankTransactionService/internal/core/service/message"
	"TrueBankTransactionService/metrics"
	"TrueBankTransactionService/pkg/models"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
	"time"
)

func GetMessageRemittance(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
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

		newRemittance := models.RemittanceHistory{
			Username:         messageResult.Username,
			SenderСardNumber: messageResult.SenderСardNumber,
			GetterCardNumber: messageResult.GetterCardNumber,
			Sum:              messageResult.Sum,
			DataCreate:       time.Now(),
		}

		metrics.KafkaMessagesOut.Inc()

		service.CreateRemittance(newRemittance)
	}
}
