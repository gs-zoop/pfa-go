package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jgsouzadev/pfa-go/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (or *OrderRepository) Save(order *entity.Order) error {
	stmt, err := or.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES(?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepository) GetTotal() (int, error) {
	var total int
	err := or.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
