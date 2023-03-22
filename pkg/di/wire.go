//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/akshayur04/project-ecommerce/pkg/api"
	handler "github.com/akshayur04/project-ecommerce/pkg/api/handler"
	config "github.com/akshayur04/project-ecommerce/pkg/config"
	db "github.com/akshayur04/project-ecommerce/pkg/db"
	repository "github.com/akshayur04/project-ecommerce/pkg/repository"
	usecase "github.com/akshayur04/project-ecommerce/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		repository.NewProductRepository,
		repository.NewCartRepository,
		repository.NewOrderRepository,
		repository.NewPaymentRepository,
		repository.NewCouponRepository,
		usecase.NewUserUseCase,
		usecase.NewOtpUseCase,
		usecase.NewAdminUsecase,
		usecase.NewFindIdUseCase,
		usecase.NewProductUsecase,
		usecase.NewOrderUseCase,
		usecase.NewCartUsecase,
		usecase.NewPaymentuseCase,
		usecase.NewCouponUsecase,
		handler.NewUserHandler,
		handler.NewOtpHandler,
		handler.NewAdminHandler,
		handler.NewProductHandler,
		handler.NewCartHandler,
		handler.NewOrderHandler,
		handler.NewPaymentHandler,
		handler.NewCouponHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
