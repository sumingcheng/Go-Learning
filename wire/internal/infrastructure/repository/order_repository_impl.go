package repository

import (
	"database/sql"

	"wire-demo/internal/domain/model"
	"wire-demo/internal/domain/repository"
)

type OrderRepositoryImpl struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) repository.OrderRepository {
	return &OrderRepositoryImpl{DB: db}
}

func (r *OrderRepositoryImpl) Save(order *model.Order) error {
	// 示例：插入订单到数据库
	result, err := r.DB.Exec("INSERT INTO orders (amount) VALUES (?)", order.Amount)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	order.ID = int(id)
	return nil
}

func (r *OrderRepositoryImpl) FindByID(id int) (*model.Order, error) {
	// 示例：从数据库中查询订单
	row := r.DB.QueryRow("SELECT id, amount FROM orders WHERE id = ?", id)
	var order model.Order
	err := row.Scan(&order.ID, &order.Amount)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
