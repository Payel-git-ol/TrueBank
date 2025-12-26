package topic_init

import (
	topic_generate "ApiGateway/internal/kafkaService/topic-generate"
	"log"
)

func InitTopic() {
	topics := []string{"register", "server"}
	for _, t := range topics {
		if err := topic_generate.CreateTopic("localhost:9092", t); err != nil {
			log.Printf("topic %s already exists or failed: %v", t, err)
		}
	}
}
