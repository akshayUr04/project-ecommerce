package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase services.ProductUsecase
}

func NewProductHandler(productUsecase services.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

// CreateCategory
// @Summary Create new product category
// @ID create-category
// @Description Admin can create new category from admin panel
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_name body helperStruct.Category true "New category name"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/category/add [post]
func (cr *ProductHandler) CreateCategory(c *gin.Context) {
	var category helperStruct.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	NewCategoery, err := cr.productUsecase.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't creat category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Created",
		Data:       NewCategoery,
		Errors:     nil,
	})
}

// UpdateCategory
// @Summary Admin can update category details
// @ID update-category
// @Description Admin can update category details
// @Tags Product Category
// @Accept json
// @Produce json
// @Param id path string true "ID of the Category to be updated"
// @Param category_details body helperStruct.Category true "category info"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/category/update/{id} [patch]
func (cr *ProductHandler) UpdatCategory(c *gin.Context) {
	var category helperStruct.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	updatedCategory, err := cr.productUsecase.UpdatCategory(category, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't update category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Updated",
		Data:       updatedCategory,
		Errors:     nil,
	})
}

// DeleteCategory
// @Summary Admin can delete a category
// @ID delete-category
// @Description Admin can delete a category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @ Router /admin/category/delete/{id} [delete]
func (cr *ProductHandler) DeleteCategory(c *gin.Context) {
	parmasId := c.Param("id")
	id, err := strconv.Atoi(parmasId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.productUsecase.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't dlete category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category deleted",
		Data:       nil,
		Errors:     nil,
	})

}

// ListAllCategories
// @Summary View all available categories
// @ID view-all-categories
// @Description Admin, users and unregistered users can see all the available categories
// @Tags Product Category
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/category/listall [get]
func (cr *ProductHandler) ListCategories(c *gin.Context) {
	categories, err := cr.productUsecase.ListCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Ctegories are",
		Data:       categories,
		Errors:     nil,
	})
}

// FindCategoryByID
// @Summary Fetch details of a specific category using category id
// @ID find-category-by-id
// @Description Users and admins can fetch details of a specific category using id
// @Tags Product Category
// @Accept json
// @Produce json
// @Param id path string true "category id"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/category/find/{id} [get]
func (cr *ProductHandler) DisplayCategory(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	category, err := cr.productUsecase.DisplayCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product is",
		Data:       category,
		Errors:     nil,
	})
}

//-----------------------------product-----------------

// CreateProduct
// @Summary Admin can create new product listings
// @ID create-product
// @Description Admins can create new product listings
// @Tags Product
// @Accept json
// @Produce json
// @Param new_product_details body helperStruct.Product true "new product details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/product/add [post]
func (cr *ProductHandler) AddProduct(c *gin.Context) {
	var product helperStruct.Product
	err := c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	newProduct, err := cr.productUsecase.AddProduct(product)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't add product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Added",
		Data:       newProduct,
		Errors:     nil,
	})

}

// UpdateProduct
// @Summary Admin can update product details
// @ID update-product
// @Description This endpoint allows an admin user to update a product's details.
// @Tags Product
// @Accept json
// @Produce json
// @Param Id path string true "ID of the product to be updated"
// @Param updated_product_details body  helperStruct.Product true "Updated product details"
// @Success 202 {object} response.Response "Successfully updated product"
// @Failure 400 {object} response.Response "Unable to update product"
// @Router /admin/product/update/{id} [put]
func (cr *ProductHandler) UpdateProduct(c *gin.Context) {
	var product helperStruct.Product
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant bind body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatedProduct, err := cr.productUsecase.UpdateProduct(id, product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant update product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, response.Response{
		StatusCode: 400,
		Message:    "Cant find id",
		Data:       updatedProduct,
		Errors:     nil,
	})

}

// DeleteProduct
// @Summary Deletes a product by ID
// @ID delete-product
// @Description This endpoint allows an admin  to delete a product by ID.
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID to delete"
// @Success 200 {object} response.Response "Successfully deleted product"
// @Failure 400 {object} response.Response "Invalid product ID"
// @Router /admin/product/delete/{id} [delete]
func (cr *ProductHandler) DeleteProduct(c *gin.Context) {

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.productUsecase.DeleteProduct(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't delete product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product deleted",
		Data:       nil,
		Errors:     nil,
	})
}

// ViewAllProducts
// @Summary Admins and users can see all available products
// @ID admin-view-all-products
// @Description Admins and users can ses all available products
// @Tags Product
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items to retrieve per page"
// @Param query query string false "Search query string"
// @Param filter query string false "Filter criteria for the products"
// @Param sort_by query string false "Sorting criteria for the products"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/product/listall [get]

// ViewAllProducts
// @Summary Admins and users can see all available products
// @ID user-view-all-products
// @Description Admins and users can ses all available products
// @Tags Product
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items to retrieve per page"
// @Param query query string false "Search query string"
// @Param filter query string false "Filter criteria for the product items"
// @Param sort_by query string false "Sorting criteria for the product items"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/products/listallproduct/ [get]
func (cr *ProductHandler) ListAllProduct(c *gin.Context) {

	var viewProduct helperStruct.QueryParams

	viewProduct.Page, _ = strconv.Atoi(c.Query("page"))
	viewProduct.Limit, _ = strconv.Atoi(c.Query("limit"))
	viewProduct.Query = c.Query("query")
	viewProduct.Filter = c.Query("filter")
	viewProduct.SortBy = c.Query("sort_by")
	viewProduct.SortDesc, _ = strconv.ParseBool(c.Query("sort_desc"))

	fmt.Println(viewProduct)

	products, err := cr.productUsecase.ListAllProduct(viewProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       products,
		Errors:     nil,
	})
}

// FindProductByID
// @Summary Admins and users can see products with product id
// @ID find-product-by-id
// @Description Admins and users can see products with product id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @ Router /admin/product/show/{id} [get]
func (cr *ProductHandler) ShowProduct(c *gin.Context) {

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	product, err := cr.productUsecase.ShowProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       product,
		Errors:     nil,
	})
}

//----------ProductItem--------

// CreateProductItem
// @Summary Creates a new product item
// @ID create-product-item
// @Description This endpoint allows an admin user to create a new product item.
// @Tags Product Item
// @Accept json
// @Produce json
// @Param product_item body helperStruct.ProductItem true "Product item details"
// @Success 200 {object} response.Response "Successfully added new product item"
// @Failure 400 {object} response.Response "Failed to add new product item"
// @Router /admin/product-item/add/ [post]
func (cr *ProductHandler) AddProductItem(c *gin.Context) {
	var productItem helperStruct.ProductItem
	err := c.Bind(&productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	newProductItem, err := cr.productUsecase.AddProductItem(productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant create",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product created",
		Data:       newProductItem,
		Errors:     nil,
	})
}

// UpdateProductItem updates a product item in the database.
// @Summary Update a product item
// @ID update-product-item
// @Description Update an existing product item with new information.
// @Tags Product Item
// @Accept json
// @Produce json
// @Param Id path string true "ID of the productitem to be updated"
// @Param product_item body helperStruct.ProductItem true "Product item information to update"
// @Success 202 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/product-item/update/{id} [patch]
func (cr *ProductHandler) UpdateProductItem(c *gin.Context) {
	var productItem helperStruct.ProductItem
	err := c.Bind(&productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	updatedItem, err := cr.productUsecase.UpdateProductItem(id, productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't update productitem",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "productitem updated",
		Data:       updatedItem,
		Errors:     nil,
	})
}

// DeleteProductItem
// @Summary Deletes a product item from the system
// @ID delete-product-item
// @Description Deletes a product item from the system
// @Tags Product Item
// @Accept json
// @Produce json
// @Param id path string true "ID of the product item to be deleted"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/product-item/delete/{id} [delete]
func (cr *ProductHandler) DeleteProductItem(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.productUsecase.DeleteProductItem(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't delete item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "item deleted",
		Data:       nil,
		Errors:     nil,
	})
}

// ViewAllProductItems
// @Summary Handler function to view all product items
// @ID admin-view-all-product-items
// @Description view all product items
// @Tags Product Item
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items to retrieve per page"
// @Param query query string false "Search query string"
// @Param filter query string false "Filter criteria for the product items"
// @Param sort_by query string false "Sorting criteria for the product items"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/product-item/listall [get]

// ViewAllProductItems for user
// @Summary Handler function to view all product items
// @ID user-view-all-product-items
// @Description view all product items for user
// @Tags Product Item
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items to retrieve per page"
// @Param query query string false "Search query string"
// @Param filter query string false "Filter criteria for the product items"
// @Param sort_by query string false "Sorting criteria for the product items"
// @Param sort_desc query bool false "Sorting in descending order"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/products/listallproductItems/ [get]
func (cr *ProductHandler) DisaplyaAllProductItems(c *gin.Context) {
	var viewProductaItem helperStruct.QueryParams

	viewProductaItem.Page, _ = strconv.Atoi(c.Query("page"))
	viewProductaItem.Limit, _ = strconv.Atoi(c.Query("limit"))
	viewProductaItem.Query = c.Query("query")
	viewProductaItem.Filter = c.Query("filter")
	viewProductaItem.SortBy = c.Query("sort_by")
	viewProductaItem.SortDesc, _ = strconv.ParseBool(c.Query("sort_desc"))

	productItems, err := cr.productUsecase.DisaplyaAllProductItems(viewProductaItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't disaply items",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product items are",
		Data:       productItems,
		Errors:     nil,
	})
}

// FindProductItemByID
// @Summary Retrieve a product item by ID
// @ID find-product-item-by-id
// @Description Retrieve a product item by its ID
// @Tags Product Item
// @Accept json
// @Produce json
// @Param id path string true "Product item ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/product-item/show/{id} [get]
func (cr *ProductHandler) DisaplyProductItem(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	productItem, err := cr.productUsecase.DisaplyProductItem(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       productItem,
		Errors:     nil,
	})

}

// UploadImage
// @Summary Handles the uploading of images for a given product ID
// @ID upload-image
// @Description Upload one or more images for a given product ID
// @Tags Product Item
// @Accept multipart/form-data
// @Param id path int true "Product ID"
// @Param images formData file true "Product image(s)"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /product-item/uploadimage/{id} [post]
func (cr *ProductHandler) UploadImage(c *gin.Context) {

	id := c.Param("id")
	productId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find product id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	// Multipart form
	form, _ := c.MultipartForm()
	fmt.Println(form)

	files := form.File["images"]

	fmt.Println(files)

	for _, file := range files {
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "asset/uploads/"+file.Filename)

		err := cr.productUsecase.UploadImage(file.Filename, productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "cant upload images",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}
}
