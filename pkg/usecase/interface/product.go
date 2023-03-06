package interfaces

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
)

type ProductUsecase interface {
	CreateCategory(category helperStruct.Category) (response.Category, error)
	UpdatCategory(category helperStruct.Category, id int) (response.Category, error)
	DeleteCategory(id int) error
	ListCategories() ([]response.Category, error)
	DisplayCategory(id int) (response.Category, error)
	AddProduct(product helperStruct.Product) (response.Product, error)
	UpdateProduct(id int, product helperStruct.Product) (response.Product, error)
	DeleteProduct(id int) error
	AddProductItem(productItem helperStruct.ProductItem) (response.ProductItem, error)
	UpdateProductItem(id int, productItem helperStruct.ProductItem) (response.ProductItem, error)
}
