package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUsecase   services.CartUsecase
	findIdUseCase services.FindIdUseCase
}

func NewCartHandler(cartUsecase services.CartUsecase, findIdUseCase services.FindIdUseCase) *CartHandler {
	return &CartHandler{
		cartUsecase:   cartUsecase,
		findIdUseCase: findIdUseCase,
	}
}

func (cr *CartHandler) AddToCart(c *gin.Context) {
	cookie, err := c.Cookie("UserAuth")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find cookie",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userId, err := cr.findIdUseCase.FindId(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	productId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.cartUsecase.AddToCart(productId, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant add product into cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product added into cart",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *CartHandler) RemoveFromCart(c *gin.Context) {
	cookie, err := c.Cookie("UserAuth")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find cookie",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userId, err := cr.findIdUseCase.FindId(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	productId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.cartUsecase.RemoveFromCart(userId, productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant remove from cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "item removed from cart",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *CartHandler) ListCart(c *gin.Context) {
	cookie, err := c.Cookie("UserAuth")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find cookie",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userId, err := cr.findIdUseCase.FindId(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	cart, err := cr.cartUsecase.ListCart(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "cart items are",
		Data:       cart,
		Errors:     nil,
	})
}

func (cr *AdminHandler) UploadImage(c *gin.Context) {

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

		err := cr.adminUseCase.UploadImage(file.Filename, productId)
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
