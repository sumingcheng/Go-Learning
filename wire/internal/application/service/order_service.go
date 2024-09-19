package service

import (
	"wire-demo/internal/domain/model"
	"wire-demo/internal/domain/repository"
)

type OrderService struct {
	OrderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	// 业务逻辑处理
	return s.OrderRepo.Save(order)
}

func (s *OrderService) GetOrder(id int) (*model.Order, error) {
	return s.OrderRepo.FindByID(id)
}
