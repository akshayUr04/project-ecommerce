package response

type CartItems struct {
	ProductItemId int
	Quantity      int
}

type Cart struct {
	// Products []CartItems
	ProductItemId int
	Quantity      int
	Tottal        int
}
