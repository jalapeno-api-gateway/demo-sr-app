// package kafka

// import (
// 	"context"
// 	"log"

// 	"github.com/Shopify/sarama"
// )

// type Consumer struct {
// 	reader *kafka.Reader
// }

// type Telemetry struct {
// 	Interface_Name string `json:"interface_name,omitempty"`
// }

// func NewConsumer(ctx context.Context, brokerAddress string, topic string) *Consumer {
// 	reader := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers:   []string{brokerAddress},
// 		Topic:     topic,
// 	})

// 	c := Consumer{reader: reader}
// 	return &c
// }

// func Consume(ctx context.Context, c Consumer, channel chan Telemetry) {
// 	for {
// 		m, err := c.reader.ReadMessage(ctx)
// 		if err != nil {
// 			log.Fatal("failed to read messages:", err)
// 			break
// 		}
// 		t := Telemetry{}
// 		kafka.Unmarshal(m.Value, &t)
// 		channel <- t
// 	}
// 	close(channel)
// }

// func ConsumeSingleMessage(ctx context.Context, c Consumer) []byte {
// 	m, err := c.reader.ReadMessage(ctx)
// 	if err != nil {
// 		log.Fatal("failed to write messages:", err)
// 	}
// 	return m.Value
// }