package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rodolfoHOk/fullcycle.go-intensivo/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery) 
	go rabbitmq.Consumer(ch, out)

	for msg := range out {
		println(string(msg.Body))
	}
}
