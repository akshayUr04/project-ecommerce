package repository

import (
	"fmt"

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

func (c *ProductDatabase) CreateCategory(category helperStruct.Category) (response.Category, error) {
	var newCategoery response.Category
	query := `INSERT INTO categories (name,created_at)VAlues($1,NOW())RETURNING id,name`
	err := c.DB.Raw(query, category.Name).Scan(&newCategoery).Error
	return newCategoery, err
}

func (c *ProductDatabase) UpdatCategory(category helperStruct.Category, id int) (response.Category, error) {
	var updatedCategory response.Category
	query := `UPDATE  categories SET name = $1 , updated_at =NOW() WHERE id=$2 RETURNING id,name `
	err := c.DB.Raw(query, category.Name, id).Scan(&updatedCategory).Error
	return updatedCategory, err
}

func (c *ProductDatabase) DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}

func (c *ProductDatabase) ListCategories() ([]response.Category, error) {
	var categories []response.Category
	query := `SELECT * FROM categories`
	err := c.DB.Raw(query).Scan(&categories).Error
	return categories, err
}

func (c *ProductDatabase) DisplayCategory(id int) (response.Category, error) {
	var category response.Category
	query := `SELECT * FROM categories WHERE id=$1`
	err := c.DB.Raw(query, id).Scan(&category).Error
	return category, err
}

func (c *ProductDatabase) AddProduct(product helperStruct.Product) (response.Product, error) {
	var newProduct response.Product
	var exits bool

	query1 := `select exists(select 1 from categories where id=?)`
	c.DB.Raw(query1, product.CategoryId).Scan(&exits)
	if !exits {
		return response.Product{}, fmt.Errorf("no category found")
	}

	query := `INSERT INTO products (name,description,brand,category_id,created_at)
		VALUES ($1,$2,$3,$4,NOW())
		RETURNING id,name,description,brand,category_id`
	err := c.DB.Raw(query, product.Name, product.Description, product.Brand, product.CategoryId).
		Scan(&newProduct).Error
	return newProduct, err
}

func (c *ProductDatabase) UpdateProduct(id int, product helperStruct.Product) (response.Product, error) {
	var updatedProduct response.Product
	query2 := `UPDATE products SET name=$1,description=$2,brand=$3,category_id=$4,updated_at=NOW() WHERE id=$5
		RETURNING id,name,description,brand,category_id`
	err := c.DB.Raw(query2, product.Name, product.Description, product.Brand, product.CategoryId, id).
		Scan(&updatedProduct).Error
	return updatedProduct, err
}

func (c *ProductDatabase) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}

func (c *ProductDatabase) AddProductItem(productItem helperStruct.ProductItem) (response.ProductItem, error) {
	var newProductItem response.ProductItem
	query := `INSERT INTO product_items (product_id,
		sku,
		qty_in_stock,
		imag,
		color,
		ram,
		battery,
		screen_size,
		storage,
		camera,
		price,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,NOW())
		RETURNING 
		id,
		product_id,
		sku,
		qty_in_stock,
		imag,
		color,
		ram,
		battery,
		screen_size,
		storage,
		camera,
		price`
	err := c.DB.Raw(query, productItem.Product_id,
		productItem.Sku,
		productItem.Qty,
		productItem.Imag,
		productItem.Color,
		productItem.Ram,
		productItem.Battery,
		productItem.Screen_size,
		productItem.Storage,
		productItem.Camera,
		productItem.Price).Scan(&newProductItem).Error
	return newProductItem, err
}

func (c *ProductDatabase) UpdateProductItem(id int, productItem helperStruct.ProductItem) (response.ProductItem, error) {
	var updatedItem response.ProductItem
	query := `UPDATE product_items SET 
	product_id=$1,
	sku=$2,
	qty_in_stock=$3,
	imag=$4,
	color=$5,
	ram=$6,
	battery=$7,
	screen_size=$8,
	storage=$9,
	camera=$10,
	price=$11,
	updated_at=NOW()
	RETURNING
		id,
		product_id,
		sku,
		qty_in_stock,
		imag,
		color,
		ram,
		battery,
		screen_size,
		storage,
		camera,
		price`
	err := c.DB.Raw(query,
		productItem.Product_id,
		productItem.Sku,
		productItem.Qty,
		productItem.Imag,
		productItem.Color,
		productItem.Ram,
		productItem.Battery,
		productItem.Screen_size,
		productItem.Storage,
		productItem.Camera,
		productItem.Price).Scan(&updatedItem).Error

	return updatedItem, err
}

func (c *ProductDatabase) DeleteProductItem(id int) error {
	query := `DELETE FROM product_items WHERE id=?`
	err := c.DB.Exec(query, id).Error
	return err
}

func (c *ProductDatabase) DisaplyaAllProductItems() ([]response.ProductItem, error) {
	var productItems []response.ProductItem
	query := `SELECT * FROM product_items`
	err := c.DB.Raw(query).Scan(&productItems).Error
	return productItems, err
}

func (c *ProductDatabase) DisaplyProductItem(id int) (response.ProductItem, error) {
	var productItem response.ProductItem
	query := `SELECT * FROM product_items WHERE  id=?`
	err := c.DB.Raw(query, id).Scan(&productItem).Error
	return productItem, err
}

func (c *ProductDatabase) ListAllProduct() ([]response.Product, error) {
	var products []response.Product
	query := `SELECT * FROM products `
	err := c.DB.Raw(query).Scan(&products).Error
	return products, err
}
