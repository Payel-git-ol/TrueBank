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

func GetTransaction(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "create-server",
		GroupID: "get-server",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		messageResult, err := message.ProcessMessageNewTransaction(msg.Value)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(messageResult)

		newTransaction := models.HistoryTransaction{
			Username:        messageResult.Username,
			NameTransaction: messageResult.NameTransaction,
			Sum:             messageResult.Sum,
			NumberCard:      messageResult.NumberCard,
			DateCreated:     time.Now(),
		}

		metrics.KafkaMessagesOut.Inc()

		service.CreateTransaction(newTransaction)
	}
}
