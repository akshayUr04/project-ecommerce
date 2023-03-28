package interfaces

import (
	"github.com/akshayur04/project-ecommerce/pkg/domain"
)

type OrderRepository interface {
	OrderAll(id, paymentTypeId int) (domain.Orders, error)
	UserCancelOrder(orderId, userId int) error
	ListOrder(userId, orderId int) (domain.Orders, error)
	ListAllOrders(userId int) ([]domain.Orders, error)
	ReturnOrder(userId, orderId int) (int, error)
}
