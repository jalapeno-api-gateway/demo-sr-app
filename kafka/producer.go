package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(ctx context.Context, brokerAddress string, topic string) *Producer {
	writer := &kafka.Writer{
		Addr: kafka.TCP(brokerAddress),
		Topic: topic,
	}

	p := Producer{writer: writer}
	return &p
}

func Produce(ctx context.Context, p Producer, message []byte) {
	err := p.writer.WriteMessages(ctx, 
		kafka.Message{
			Value: message,
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}
