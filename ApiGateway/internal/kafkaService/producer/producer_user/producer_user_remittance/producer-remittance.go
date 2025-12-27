package producer_user_remittance

import (
	"ApiGateway/pkg/models/remittance"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func SendMessageRemittance(topic string, data remittance.RemittanceTransaction) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = w.WriteMessages(context.Background(), kafka.Message{
		Value: jsonData,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Отправлено в топик '%s': %v\n", topic, data)
	return nil
}
