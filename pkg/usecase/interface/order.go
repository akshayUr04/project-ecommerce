package interfaces

type OrderUseCase interface {
	PlaceOrder(id int) error
}
