package producer

import (
	"ApiGateway/pkg/models"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

func SendTransaction(topic string, data models.Transaction) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling transaction %v", err)
		return err
	}

	err = w.WriteMessages(context.Background(), kafka.Message{
		Value: jsonData,
	})
	if err != nil {
		log.Fatalf("Error sending message %v", err)
		return err
	}

	return nil
}
