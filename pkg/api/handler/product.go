package handler

import (
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

func (cr *ProductHandler) ListAllProduct(c *gin.Context) {
	products, err := cr.productUsecase.ListAllProduct()
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

func (cr *ProductHandler) DisaplyaAllProductItems(c *gin.Context) {
	productItems, err := cr.productUsecase.DisaplyaAllProductItems()

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
