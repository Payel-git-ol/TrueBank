package producer_transaction

import (
	"ApiGateway/pkg/models/transaction/request"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

func SendMessageReplenishment(topic string, replenishment request.Replenishment) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(replenishment)
	if err != nil {
		log.Println(err)
		return err
	}

	msg := w.WriteMessages(context.Background(), kafka.Message{Value: jsonData})
	log.Println(msg)
	return nil
}
