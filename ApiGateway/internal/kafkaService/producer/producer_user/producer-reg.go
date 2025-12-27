package producer_user

import (
	"ApiGateway/pkg/models/user"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func SendMessageInRegistretion(topic string, data user.User) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{Value: jsonData},
	)
	if err != nil {
		panic(err)
	}

	//metrics.KafkaMessagesOut.Inc()

	fmt.Printf("Отправлено в топик '%s': %v\n", topic, string(jsonData))
	return nil
}
