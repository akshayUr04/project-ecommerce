package response

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id          int
	Name        string
	Description string
	Brand       string
	CategoryId  string
}

type ProductItem struct {
	Id           uint
	Product_id   uint
	Sku          string
	Qty_in_stock int
	Color        string
	Ram          int
	Battery      int
	Screen_size  float64
	Storage      int
	Camera       int
	Price        int
	Imag         string
}
