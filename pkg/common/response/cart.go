package response

type CartItems struct {
	ProductItemId int
	Quantity      int
}

type Cart struct {
	// Products []CartItems
	Sku      string
	Color    string
	Price    int
	Quantity int
	Tottal   int
}
