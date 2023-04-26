package usecase

import (
	"github.com/akshayur04/project-ecommerce/pkg/domain"
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

func (c *OrderUseCase) OrderAll(id, paymentTypeId int) (domain.Orders, error) {
	order, err := c.orderRepo.OrderAll(id, paymentTypeId)
	return order, err
}

func (c *OrderUseCase) UserCancelOrder(orderId, userId int) error {
	err := c.orderRepo.UserCancelOrder(orderId, userId)
	return err
}

func (c *OrderUseCase) ListOrder(userId, orderId int) (domain.Orders, error) {
	order, err := c.orderRepo.ListOrder(userId, orderId)
	return order, err
}

func (c *OrderUseCase) ListAllOrders(userId int) ([]domain.Orders, error) {
	orders, err := c.orderRepo.ListAllOrders(userId)
	return orders, err
}

func (c *OrderUseCase) ReturnOrder(userId, orderId int) (int, error) {
	returnAmount, err := c.orderRepo.ReturnOrder(userId, orderId)
	return returnAmount, err
}

func (c *OrderUseCase) UpdateOrder(orderId int) error {
	err := c.orderRepo.UpdateOrder(orderId)
	return err
}
