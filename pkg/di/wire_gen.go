// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/akshayur04/project-ecommerce/pkg/api"
	"github.com/akshayur04/project-ecommerce/pkg/api/handler"
	"github.com/akshayur04/project-ecommerce/pkg/config"
	"github.com/akshayur04/project-ecommerce/pkg/db"
	"github.com/akshayur04/project-ecommerce/pkg/repository"
	"github.com/akshayur04/project-ecommerce/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	otpUseCase := usecase.NewOtpUseCase(cfg)
	otpHandler := handler.NewOtpHandler(cfg, otpUseCase, userUseCase)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUsecase := usecase.NewAdminUsecase(adminRepository)
	findIdUseCase := usecase.NewFindIdUseCase()
	adminHandler := handler.NewAdminHandler(adminUsecase, findIdUseCase)
	productRepository := repository.NewProductRepository(gormDB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)
	serverHTTP := http.NewServerHTTP(userHandler, otpHandler, adminHandler, productHandler)
	return serverHTTP, nil
}
