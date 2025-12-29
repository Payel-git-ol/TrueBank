package producer

import (
	"TrueBankTransactionService/metrics"
	"TrueBankTransactionService/pkg/models/respons"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

func SendMessageRemittance(topic string, data respons.ResultRemittance) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
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

	metrics.KafkaMessagesIn.Inc()

	log.Println("Send in topic: ", topic, "Remittance: ", data)
}
