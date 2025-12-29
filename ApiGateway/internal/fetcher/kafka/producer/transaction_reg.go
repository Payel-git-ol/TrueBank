package producer

import (
	"ApiGateway/metrics"
	"ApiGateway/pkg/model"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

func SendTransactionReg(topic string, data model.RegTransaction) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling server %v", err)
		return err
	}

	err = w.WriteMessages(context.Background(), kafka.Message{
		Value: jsonData,
	})
	if err != nil {
		log.Fatalf("Error sending message %v", err)
		return err
	}

	metrics.KafkaMessagesIn.Inc()

	return nil
}
