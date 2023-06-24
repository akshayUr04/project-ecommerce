package repository

import (
	"fmt"
	"strings"

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
	query := `INSERT INTO categories (category_name,created_at)VAlues($1,NOW())RETURNING id,category_name`
	err := c.DB.Raw(query, category.Name).Scan(&newCategoery).Error
	return newCategoery, err
}

func (c *ProductDatabase) UpdatCategory(category helperStruct.Category, id int) (response.Category, error) {
	var updatedCategory response.Category
	query := `UPDATE  categories SET category_name = $1 , updated_at =NOW() WHERE EXISTS(SELECT 1 FROM categories WHERE id=$2) RETURNING id,category_name `
	err := c.DB.Raw(query, category.Name, id).Scan(&updatedCategory).Error
	if err != nil {
		return response.Category{}, err
	}
	if updatedCategory.Id == 0 {
		return response.Category{}, fmt.Errorf("no such category to update")
	}
	return updatedCategory, nil
}

func (c *ProductDatabase) DeleteCategory(id int) error {
	var exits bool

	query1 := `select exists(select 1 from categories where id=?)`
	c.DB.Raw(query1, id).Scan(&exits)
	if !exits {
		return fmt.Errorf("no category found")
	}
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
	var exits bool

	query1 := `select exists(select 1 from categories where id=?)`
	c.DB.Raw(query1, id).Scan(&exits)
	if !exits {
		return response.Category{}, fmt.Errorf("no category found")
	}
	query := `SELECT * FROM categories WHERE id=$1`
	err := c.DB.Raw(query, id).Scan(&category).Error
	if err != nil {
		return response.Category{}, err
	}
	if category.Id == 0 {
		return response.Category{}, fmt.Errorf("no such category")
	}
	return category, nil
}

func (c *ProductDatabase) AddProduct(product helperStruct.Product) (response.Product, error) {
	var newProduct response.Product
	var exits bool

	query1 := `select exists(select 1 from categories where id=?)`
	c.DB.Raw(query1, product.CategoryId).Scan(&exits)
	if !exits {
		return response.Product{}, fmt.Errorf("no category found")
	}

	// query := `INSERT INTO products (product_name, description, brand, category_id, created_at)
	// 	VALUES ($1, $2, $3, $4, NOW())
	// 	RETURNING p.id, p.product_name AS name, p.description, p.brand, c.category_name AS category
	// 	FROM products p
	// 	JOIN categories c
	// 	ON p.category_id = c.id`

	query := `INSERT INTO products (product_name,description,brand,category_id,created_at)
		VALUES ($1,$2,$3,$4,NOW())
		RETURNING id,product_name AS name,description,brand,category_id`
	err := c.DB.Raw(query, product.Name, product.Description, product.Brand, product.CategoryId).
		Scan(&newProduct).Error
	return newProduct, err
}

func (c *ProductDatabase) UpdateProduct(id int, product helperStruct.Product) (response.Product, error) {
	var updatedProduct response.Product
	query2 := `UPDATE products SET product_name=$1,description=$2,brand=$3,category_id=$4,updated_at=NOW() WHERE id=$5
		RETURNING id,product_name,description,brand,category_id`
	err := c.DB.Raw(query2, product.Name, product.Description, product.Brand, product.CategoryId, id).
		Scan(&updatedProduct).Error
	if err != nil {
		return response.Product{}, err
	}
	if updatedProduct.Id == 0 {
		return response.Product{}, fmt.Errorf("there is no such product")
	}
	return updatedProduct, nil
}

func (c *ProductDatabase) DeleteProduct(id int) error {
	var exists bool
	isExists := `SELECT EXISTS (SELECT 1 FROM products WHERE id=$1)`
	c.DB.Raw(isExists, id).Scan(&exists)
	if !exists {
		return fmt.Errorf("there is no such product to delete")
	}
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
	var exists bool
	isExists := `SELECT EXISTS (SELECT 1 FROM product_items WHERE id=$1)`
	c.DB.Raw(isExists, id).Scan(&exists)
	if !exists {
		return response.ProductItem{}, fmt.Errorf("there is no such product to update")
	}

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
	WHERE id=$12
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
		productItem.Price,
		id).Scan(&updatedItem).Error

	return updatedItem, err
}

func (c *ProductDatabase) DeleteProductItem(id int) error {
	var exists bool
	isExists := `SELECT EXISTS (SELECT 1 FROM product_items WHERE id=$1)`
	c.DB.Raw(isExists, id).Scan(&exists)
	if !exists {
		return fmt.Errorf("there is no such product to delete")
	}
	query := `DELETE FROM product_items WHERE id=?`
	err := c.DB.Exec(query, id).Error
	return err
}

func (c *ProductDatabase) DisaplyaAllProductItems(queryParams helperStruct.QueryParams) ([]response.ProductItem, error) {
	var productItems []response.ProductItem
	getProductItemDetails := `SELECT p.product_name,
		p.description,
		p.brand,
		c.category_name, 
		pi.*
		FROM products p 
		JOIN categories c ON p.category_id=c.id 
		JOIN product_items pi ON p.id=pi.product_id`

	if queryParams.Query != "" && queryParams.Filter != "" {
		getProductItemDetails = fmt.Sprintf("%s WHERE LOWER(%s) LIKE '%%%s%%'", getProductItemDetails, queryParams.Filter, strings.ToLower(queryParams.Query))
	}

	if queryParams.SortBy != "" {
		if queryParams.SortDesc {
			getProductItemDetails = fmt.Sprintf("%s ORDER BY %s DESC", getProductItemDetails, queryParams.SortBy)
		} else {
			getProductItemDetails = fmt.Sprintf("%s ORDER BY %s ASC", getProductItemDetails, queryParams.SortBy)
		}
	} else {
		getProductItemDetails = fmt.Sprintf("%s ORDER BY p.created_at DESC", getProductItemDetails)
	}
	//to set the page number and the qty that need to display in a single responce
	if queryParams.Limit != 0 && queryParams.Page != 0 {
		getProductItemDetails = fmt.Sprintf("%s LIMIT %d OFFSET %d", getProductItemDetails, queryParams.Limit, (queryParams.Page-1)*queryParams.Limit)
	}
	if queryParams.Limit == 0 || queryParams.Page == 0 {
		getProductItemDetails = fmt.Sprintf("%s LIMIT 10 OFFSET 0", getProductItemDetails)
	}

	err := c.DB.Raw(getProductItemDetails).Scan(&productItems).Error
	return productItems, err
}

func (c *ProductDatabase) DisaplyProductItem(id int) (response.ProductItem, error) {
	var productItem response.ProductItem
	query := `SELECT p.product_name,
	p.description,
	p.brand,
	c.category_name, 
	pi.*
	FROM products p 
	JOIN categories c ON p.category_id=c.id 
	JOIN product_items pi ON p.id=pi.product_id 
	WHERE pi.id=$1`
	err := c.DB.Raw(query, id).Scan(&productItem).Error
	if err != nil {
		return response.ProductItem{}, err
	}
	if productItem.Id == 0 {
		return response.ProductItem{}, fmt.Errorf("there is no such product item")
	}
	getImages := `SELECT file_name FROM images WHERE product_item_id=$1`
	err = c.DB.Raw(getImages, id).Scan(&productItem.Image).Error
	if err != nil {
		return response.ProductItem{}, err
	}
	return productItem, nil
}

func (c *ProductDatabase) ListAllProduct(queryParams helperStruct.QueryParams) ([]response.Product, error) {
	var products []response.Product
	getProductDetails := `SELECT p.product_name AS name,
		p.description,
		p.brand,
		c.category_name
		 FROM products p JOIN categories c ON p.category_id=c.id `
	if queryParams.Query != "" && queryParams.Filter != "" {
		getProductDetails = fmt.Sprintf("%s WHERE LOWER(%s) LIKE '%%%s%%'", getProductDetails, queryParams.Filter, strings.ToLower(queryParams.Query))
	}

	if queryParams.SortBy != "" {
		if queryParams.SortDesc {
			getProductDetails = fmt.Sprintf("%s ORDER BY %s DESC", getProductDetails, queryParams.SortBy)
		} else {
			getProductDetails = fmt.Sprintf("%s ORDER BY %s ASC", getProductDetails, queryParams.SortBy)
		}
	} else {
		getProductDetails = fmt.Sprintf("%s ORDER BY p.created_at DESC", getProductDetails)
	}
	//to set the page number and the qty that need to display in a single responce
	if queryParams.Limit != 0 && queryParams.Page != 0 {
		getProductDetails = fmt.Sprintf("%s LIMIT %d OFFSET %d", getProductDetails, queryParams.Limit, (queryParams.Page-1)*queryParams.Limit)
	}
	if queryParams.Limit == 0 || queryParams.Page == 0 {
		getProductDetails = fmt.Sprintf("%s LIMIT 10 OFFSET 0", getProductDetails)
	}

	fmt.Println(getProductDetails)
	err := c.DB.Raw(getProductDetails).Scan(&products).Error
	if err != nil {
		return []response.Product{}, err
	}
	return products, nil
}

func (c *ProductDatabase) ShowProduct(id int) (response.Product, error) {
	var product response.Product
	query := `SELECT p.product_name,p.description,p.brand,c.category_name FROM products p 
		JOIN categories c ON p.category_id=c.id WHERE p.id=$1`
	err := c.DB.Raw(query, id).Scan(&product).Error
	if err != nil {
		return response.Product{}, err
	}
	if product.Id == 0 {
		return response.Product{}, fmt.Errorf("there is no such product")
	}
	return product, err
}

func (c *ProductDatabase) UploadImage(filepath string, productId int) error {
	uploadImage := `INSERT INTO images (product_item_id,file_name)VALUES($1,$2)`
	err := c.DB.Exec(uploadImage, productId, filepath).Error
	return err
}
