package producer

import (
	"TrueBankTransactionService/pkg/models/requestModels"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

func SendMessageRemittance(topic string, data requestModels.ResultRemittance) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	w.WriteMessages(context.Background(), kafka.Message{
		Value: jsonData,
	})

	log.Println("Send in topic: ", topic, "Remittance: ", data)
}
