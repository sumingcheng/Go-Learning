//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"wire-demo/internal/application/service"
	"wire-demo/internal/infrastructure/db"
	infraRepo "wire-demo/internal/infrastructure/repository"
	"wire-demo/internal/interfaces/handler"
)

func InitializeOrderHandler(dsn string) (*handler.OrderHandler, error) {
	wire.Build(
		db.NewDatabase,
		infraRepo.NewOrderRepository,
		//wire.Bind(new(repository.OrderRepository), new(*infraRepo.OrderRepositoryImpl)),
		service.NewOrderService,
		handler.NewOrderHandler,
	)
	return &handler.OrderHandler{}, nil
}
