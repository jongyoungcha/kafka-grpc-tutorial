package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jongyoungcha/kafka-grpc-tutorial/protocols"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const (
	topic     = "foobar"
	partition = 0
)

type IProducer interface {
	PubFoo(foo *protocols.MessageFoo) error
	SubBar() (*protocols.MessageBar, error)
}

type producer struct {
}

func NewProducer() IProducer {
	return &producer{}
}

func (p producer) PubFoo(foo *protocols.MessageFoo) error {
	// to produce messages
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	data, err := json.Marshal(foo)
	if err != nil {
		log.Fatal("marshalling failed: ", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: data},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	return nil
}

func (p producer) SubBar() (*protocols.MessageBar, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}

	return &protocols.MessageBar{
		Version: "",
		Index:   0,
		Bar:     "",
		Messge:  "",
	}, nil
}
