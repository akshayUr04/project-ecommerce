package response

type Category struct {
	Id           int
	CategoryName string
}

type Product struct {
	Id           int
	Name         string
	Description  string
	Brand        string
	CategoryName string
}

type ProductItem struct {
	Id           uint
	ProductName  string
	Description  string
	Brand        string
	CategoryName string
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
