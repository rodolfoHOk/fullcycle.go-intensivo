package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

	// forever := make(chan bool) // necessário para manter a aplicação rodando antes de inserirmos o servidor http
	quantityWorker := 50
	for i := 1; i <= quantityWorker; i++ {
		go worker(out, uc, i)
	}
	// <-forever // necessário para manter a aplicação rodando antes de inserirmos o servidor http

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request){
		getTotalUC := usecase.NewGetTotalUseCase(repository)
		total, err := getTotalUC.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		json.NewEncoder(w).Encode(total)
	})

	http.ListenAndServe(":8080", nil)
}

func worker(deliveryMessage <- chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryMessage {
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
		fmt.Printf("Worker %d has processed order %s\n", workerId, outputDTO.ID)
		time.Sleep(1 * time.Second)
	}
}
