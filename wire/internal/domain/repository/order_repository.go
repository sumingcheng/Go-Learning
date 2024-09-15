package repository

import "wire-demo/internal/domain/model"

type OrderRepository interface {
	Save(order *model.Order) error
	FindByID(id int) (*model.Order, error)
}
