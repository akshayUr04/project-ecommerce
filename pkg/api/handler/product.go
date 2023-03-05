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
			Message:    "can't find product",
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
