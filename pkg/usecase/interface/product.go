package interfaces

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
)

type ProductUsecase interface {
	CreateCategory(category helperStruct.Category) (response.CategoryResp, error)
	UpdatCategory(category helperStruct.Category, id int) (response.CategoryResp, error)
	DeleteCategory(id int) error
	ListCategories() ([]response.CategoryResp, error)
	DisplayCategory(id int) (response.CategoryResp, error)
}
