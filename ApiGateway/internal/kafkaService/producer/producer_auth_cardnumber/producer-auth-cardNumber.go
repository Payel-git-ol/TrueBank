package producer_auth_cardnumber

import (
	"ApiGateway/pkg/models/cardNumber"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func SendMessageAuthCardNumber(topic string, data cardNumber.AuthCardNumber) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(data)
	err = w.WriteMessages(context.Background(),
		kafka.Message{Value: jsonData},
	)
	if err != nil {
		panic(err)
		return err
	}

	//metrics.KafkaMessagesOut.Inc()

	fmt.Printf("Отправлено в топик '%s': %v\n", topic, jsonData)
	return nil
}
