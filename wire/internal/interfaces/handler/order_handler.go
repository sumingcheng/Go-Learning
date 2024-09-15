package handler

import (
	"fmt"

	"wire-demo/internal/application/service"
	"wire-demo/internal/domain/model"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

func (h *OrderHandler) HandleCreateOrder() {
	order := &model.Order{Amount: 200.0}
	err := h.OrderService.CreateOrder(order)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return
	}
	fmt.Printf("Order created successfully with ID: %d\n", order.ID)
}

func (h *OrderHandler) HandleGetOrder(id int) {
	order, err := h.OrderService.GetOrder(id)
	if err != nil {
		fmt.Println("Error getting order:", err)
		return
	}
	fmt.Printf("Order retrieved: %+v\n", order)
}
