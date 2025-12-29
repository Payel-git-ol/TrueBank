package producer

import (
	"TrueBankTransactionService/pkg/models/respons"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func SendMessageTransaction(topic string, sum float64, numberCard string, username string) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	defer w.Close()

	data := respons.ResultTransaction{
		Username:   username,
		Sum:        sum,
		CardNumber: numberCard,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = w.WriteMessages(context.Background(), kafka.Message{
		Value: jsonData,
	})

	fmt.Printf("Отправлено в топик '%s': %v\n", topic, data)
	return nil
}
