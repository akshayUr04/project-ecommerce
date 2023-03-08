package usecase

import (
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
)

type OrderUseCase struct {
	orderRepo interfaces.OrderRepository
}

func NewOrderUseCase(orderRepo interfaces.OrderRepository) services.OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
	}
}

func (c *OrderUseCase) PlaceOrder(id int) error {
	err := c.orderRepo.PlaceOrder(id)
	return err
}
