package main

import (
	"fmt"
	"github.com/jongyoungcha/kafka-grpc-tutorial/protocols"
	"testing"
)

func TestProducer(t *testing.T) {
	producer := NewProducer()

	fooMsg := &protocols.MessageFoo{
		Version: "1",
		Index:   0,
		Foo:     "foo",
		Messge:  "message!!!",
	}

	err := producer.PubFoo(fooMsg)
	if err != nil {
		t.Error("foobar: ", err)
		return
	}
}

func TestConsumer(t *testing.T) {
	consumer := NewConsumer()

	msg, err := consumer.SubFoo()
	if err != nil {
		t.Error("foobar: ", err)
		return
	}
	fmt.Println("msg: ", msg)

	barMsg := &protocols.MessageBar{
		Version: "1",
		Index:   0,
		Bar:     "bar",
		Messge:  "message!!!",
	}

	err = consumer.PubBar(barMsg)
	if err != nil {
		t.Error("pubbar: ", err)
		return
	}
}
