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

func (c *ProductUsecase) CreateCategory(category helperStruct.Category) (response.CategoryResp, error) {
	newCategory, err := c.productRepo.CreateCategory(category)
	return newCategory, err
}

func (c *ProductUsecase) UpdatCategory(category helperStruct.Category, id int) (response.CategoryResp, error) {
	updatedCategory, err := c.productRepo.UpdatCategory(category, id)
	return updatedCategory, err
}

func (c *ProductUsecase) DeleteCategory(id int) error {
	err := c.productRepo.DeleteCategory(id)
	return err
}

func (c *ProductUsecase) ListCategories() ([]response.CategoryResp, error) {
	categories, err := c.productRepo.ListCategories()
	return categories, err
}

func (c *ProductUsecase) DisplayCategory(id int) (response.CategoryResp, error) {
	category, err := c.productRepo.DisplayCategory(id)
	return category, err
}
