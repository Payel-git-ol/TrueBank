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
)

func GetRegTransaction(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "server-reg",
		GroupID: "get-reg-server",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		processMessage, err := message.ProcessMessageRegTransaction(msg.Value)
		if err != nil {
			log.Fatal(err)
		}

		newTransaction := models.ListTransaction{
			Name:                             processMessage.Name,
			Description:                      processMessage.Description,
			Company:                          processMessage.Company,
			Documents:                        processMessage.Documents,
			LinkToIndividualEntrepreneurship: processMessage.LinkToIndividualEntrepreneurship,
		}

		service.SaveRegTransaction(newTransaction)

		metrics.KafkaMessagesOut.Inc()

		fmt.Println(processMessage)
	}
}
