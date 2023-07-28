package main

import (
	"database/sql"

	"github.com/jgsouzadev/pfa-go/internal/order/infra/database"
	"github.com/jgsouzadev/pfa-go/internal/order/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderInputDTO{
		ID:    "1234",
		Price: 100.0,
		Tax:   10.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}

	println(output.FinalPrice)
}
