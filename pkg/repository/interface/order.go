package interfaces

import "github.com/akshayur04/project-ecommerce/pkg/domain"

type OrderRepository interface {
	OrderAll(id, paymentTypeId int) (domain.Orders, error)
	UserCancelOrder(orderId, userId int) error
}
