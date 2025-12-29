package producer

import (
	"ApiGateway/metrics"
	"ApiGateway/pkg/model/requests"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

func SendMessageReplenishment(topic string, replenishment requests.Replenishment) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(replenishment)
	if err != nil {
		log.Println(err)
		return err
	}

	metrics.KafkaMessagesIn.Inc()

	msg := w.WriteMessages(context.Background(), kafka.Message{Value: jsonData})
	log.Println(msg)
	return nil
}
