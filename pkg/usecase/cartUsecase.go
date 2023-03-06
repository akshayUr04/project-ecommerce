package usecase

import (
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
)

type cartUseCase struct {
	cartRepo interfaces.CartRepository
}

func NewcartUsecase(cartRepo interfaces.CartRepository) services.CartUsecase {
	return &cartUseCase{
		cartRepo: cartRepo,
	}
}
