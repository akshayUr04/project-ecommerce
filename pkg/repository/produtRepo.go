package repository

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type ProductDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &ProductDatabase{DB}
}

func (c *ProductDatabase) CreateCategory(category helperStruct.Category) (response.CategoryResp, error) {
	var newCategoery response.CategoryResp
	query := `INSERT INTO categories (name,created_at)VAlues($1,NOW())RETURNING id,name`
	err := c.DB.Raw(query, category.Name).Scan(&newCategoery).Error
	return newCategoery, err
}

func (c *ProductDatabase) UpdatCategory(category helperStruct.Category, id int) (response.CategoryResp, error) {
	var updatedCategory response.CategoryResp
	query := `UPDATE  categories SET name = $1 , updated_at =NOW() WHERE id=$2 RETURNING id,name `
	err := c.DB.Raw(query, category.Name, id).Scan(&updatedCategory).Error
	return updatedCategory, err
}

func (c *ProductDatabase) DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}

func (c *ProductDatabase) ListCategories() ([]response.CategoryResp, error) {
	var categories []response.CategoryResp
	query := `SELECT * FROM categories`
	err := c.DB.Raw(query).Scan(&categories).Error
	return categories, err
}

func (c *ProductDatabase) DisplayCategory(id int) (response.CategoryResp, error) {
	var category response.CategoryResp
	query := `SELECT * FROM categories WHERE id=$1`
	err := c.DB.Raw(query, id).Scan(&category).Error
	return category, err
}
