package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jgsouzadev/pfa-go/internal/order/infra/database"
	"github.com/jgsouzadev/pfa-go/internal/order/usecase"
	"github.com/jgsouzadev/pfa-go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	wg := sync.WaitGroup{}
	maxWorkers := 20
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, out)
	wg.Add(maxWorkers)
	for i := 0; 1 < maxWorkers; i++ {
		defer wg.Done()
		go worker(out, uc, i)
	}
	wg.Wait()
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryMessage {
		var input usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			fmt.Println(err)
		}

		input.Tax = 10.0
		_, err = uc.Execute(input)

		if err != nil {
			fmt.Println(err)
		}

		//println("WorkerId=", workerId, ", processed order=", input.ID)
		msg.Ack(false)

	}
}
