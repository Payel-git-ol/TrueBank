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
)

func GetRegTransaction(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
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

		newTransaction := dbModels.ListTransaction{
			Name:                             processMessage.Name,
			Description:                      processMessage.Description,
			Company:                          processMessage.Company,
			Documents:                        processMessage.Documents,
			LinkToIndividualEntrepreneurship: processMessage.LinkToIndividualEntrepreneurship,
		}

		service.SaveRegTransaction(newTransaction)

		fmt.Println(processMessage)
	}
}
