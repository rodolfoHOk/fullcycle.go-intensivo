package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rodolfoHOk/fullcycle.go-intensivo/internal/order/entity"
)

func Publish(ch *amqp.Channel, order entity.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func generateOrder() entity.Order {
	return entity.Order{
		ID: uuid.New().String(),
		Price: rand.Float64() * 100,
		Tax: rand.Float64() * 10,
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 10000000; i++ { // 100 to 10000000 for testing Grafana + Prometheus only
		Publish(ch, generateOrder())
		time.Sleep(300 * time.Millisecond) // add line for testing Grafana + Prometheus only
	}
}
