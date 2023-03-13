package usecase

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
)

type ProductUsecase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUsecase(productRepo interfaces.ProductRepository) services.ProductUsecase {
	return &ProductUsecase{
		productRepo: productRepo,
	}
}

func (c *ProductUsecase) CreateCategory(category helperStruct.Category) (response.Category, error) {
	newCategory, err := c.productRepo.CreateCategory(category)
	return newCategory, err
}

func (c *ProductUsecase) UpdatCategory(category helperStruct.Category, id int) (response.Category, error) {
	updatedCategory, err := c.productRepo.UpdatCategory(category, id)
	return updatedCategory, err
}

func (c *ProductUsecase) DeleteCategory(id int) error {
	err := c.productRepo.DeleteCategory(id)
	return err
}

func (c *ProductUsecase) ListCategories() ([]response.Category, error) {
	categories, err := c.productRepo.ListCategories()
	return categories, err
}

func (c *ProductUsecase) DisplayCategory(id int) (response.Category, error) {
	category, err := c.productRepo.DisplayCategory(id)
	return category, err
}

func (c *ProductUsecase) AddProduct(product helperStruct.Product) (response.Product, error) {
	newProduct, err := c.productRepo.AddProduct(product)
	return newProduct, err
}

func (c *ProductUsecase) UpdateProduct(id int, product helperStruct.Product) (response.Product, error) {
	updatedProduct, err := c.productRepo.UpdateProduct(id, product)
	return updatedProduct, err
}

func (c *ProductUsecase) DeleteProduct(id int) error {
	err := c.productRepo.DeleteProduct(id)
	return err
}

func (c *ProductUsecase) AddProductItem(productItem helperStruct.ProductItem) (response.ProductItem, error) {
	newProductItem, err := c.productRepo.AddProductItem(productItem)
	return newProductItem, err
}

func (c *ProductUsecase) UpdateProductItem(id int, productItem helperStruct.ProductItem) (response.ProductItem, error) {
	updatedItem, err := c.productRepo.UpdateProductItem(id, productItem)
	return updatedItem, err
}

func (c *ProductUsecase) DeleteProductItem(id int) error {
	err := c.productRepo.DeleteProductItem(id)
	return err
}

func (c *ProductUsecase) DisaplyaAllProductItems() ([]response.ProductItem, error) {
	productItems, err := c.productRepo.DisaplyaAllProductItems()
	return productItems, err
}

func (c *ProductUsecase) DisaplyProductItem(id int) (response.ProductItem, error) {
	productItem, err := c.productRepo.DisaplyProductItem(id)
	return productItem, err
}

func (c *ProductUsecase) ListAllProduct() ([]response.Product, error) {
	products, err := c.productRepo.ListAllProduct()
	return products, err
}

func (c *ProductUsecase) ShowProduct(id int) (response.Product, error) {
	product, err := c.productRepo.ShowProduct(id)
	return product, err
}
