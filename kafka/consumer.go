package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(ctx context.Context, brokerAddress string, topic string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddress},
		Topic:     topic,
	})

	c := Consumer{reader: reader}
	return &c
}

func Consume(ctx context.Context, c Consumer, channel chan []byte) {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Fatal("failed to read messages:", err)
			break
		}
		channel <- m.Value
	}
	close(channel)
}

func ConsumeSingleMessage(ctx context.Context, c Consumer) []byte {
	m, err := c.reader.ReadMessage(ctx)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	return m.Value
}