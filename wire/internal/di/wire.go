//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"wire/internal/application/service"
	"wire/internal/infrastructure/db"
	infraRepo "wire/internal/infrastructure/repository"
	"wire/internal/interfaces/handler"
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
