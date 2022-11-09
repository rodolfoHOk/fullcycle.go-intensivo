package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rodolfoHOk/fullcycle.go-intensivo/internal/order/infra/database"
	"github.com/rodolfoHOk/fullcycle.go-intensivo/internal/order/usecase"
	"github.com/rodolfoHOk/fullcycle.go-intensivo/pkg/rabbitmq"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:mysqluser@tcp(localhost:3306)/ordersdb?allowNativePasswords=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(*repository)
	
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery) 
	go rabbitmq.Consumer(ch, out)

	for msg := range out {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}
		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println(outputDTO)
		// time.Sleep(500 * time.Millisecond) // add line for testing Grafana + Prometheus only
	}
}
