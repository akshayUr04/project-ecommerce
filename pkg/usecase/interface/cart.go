package interfaces

import "github.com/akshayur04/project-ecommerce/pkg/common/response"

type CartUsecase interface {
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(userId, productId int) error
	ListCart(userId int) ([]response.Cart, error)
}
