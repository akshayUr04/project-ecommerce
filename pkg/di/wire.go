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
		usecase.NewUserUseCase,
		usecase.NewOtpUseCase,
		usecase.NewAdminUsecase,
		usecase.NewFindIdUseCase,
		handler.NewUserHandler,
		handler.NewOtpHandler,
		handler.NewAdminHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
