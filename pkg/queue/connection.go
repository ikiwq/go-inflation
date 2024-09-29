package queue

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func InitConnection(networkType string, address string, topic string, partition int) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), networkType, address, topic, partition)
	if err != nil {
		log.Fatal("failed to load leader:", err)
	}

	return conn
}

func InitReader(brokers []string, topic string, maxWait time.Duration, minBytes int, maxBytes int) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MaxWait:  maxWait,
		MinBytes: minBytes,
		MaxBytes: maxBytes,
	})
}
