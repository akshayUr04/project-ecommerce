package interfaces

import "github.com/akshayur04/project-ecommerce/pkg/domain"

type OrderUseCase interface {
	OrderAll(id, paymentTypeId int) (domain.Orders, error)
}
