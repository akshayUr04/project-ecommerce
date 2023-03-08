package interfaces

type OrderRepository interface {
	PlaceOrder(id int) error
}
