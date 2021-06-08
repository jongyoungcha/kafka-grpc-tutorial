package main

import (
	"context"
	"fmt"
	"github.com/jongyoungcha/kafka-grpc-tutorial/protocols"
	"github.com/segmentio/kafka-go"
	"log"
)

type IConsumer interface {
	PubBar(foo *protocols.MessageBar) error
	SubFoo() (*protocols.MessageFoo, error)
}

type consumer struct {
}

func NewConsumer() IConsumer {
	return &consumer{}
}

func (c consumer) PubBar(foo *protocols.MessageBar) error {
	return nil
}

func (c consumer) SubFoo() (*protocols.MessageFoo, error) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "foobar",
		GroupID:   "test-consumer-group",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(42)
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

	return nil, nil
}
