package interfaces

type CartUsecase interface {
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(userId, productId int) error
}
