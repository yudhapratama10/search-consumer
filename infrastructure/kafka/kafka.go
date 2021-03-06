package kafka

import (
	"github.com/segmentio/kafka-go"
)

func GetClient() *kafka.Reader {
	brokers := []string{
		"localhost:9092",
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   "search",
		GroupID: "test",
	})

	return reader
}
