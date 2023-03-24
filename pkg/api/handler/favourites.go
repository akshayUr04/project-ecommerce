package handler

import (
	"net/http"
	"strconv"

	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type FavouriteHandler struct {
	favouritesUsecase services.FavouritesUsecase
	findIdUseCase     services.FindIdUseCase
}

func NewFavouritesHandler(favouritesUsecase services.FavouritesUsecase, findIdUseCase services.FindIdUseCase) *FavouriteHandler {
	return &FavouriteHandler{
		favouritesUsecase: favouritesUsecase,
		findIdUseCase:     findIdUseCase,
	}
}

func (cr *FavouriteHandler) AddToFavourites(c *gin.Context) {
	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find the productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
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
	err = cr.favouritesUsecase.AddToFavourites(productId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant add product to favourites",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product added to favourites",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *FavouriteHandler) RemoveFromFav(c *gin.Context) {
	id := c.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find the productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
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
	err = cr.favouritesUsecase.RemoveFromFav(userId, productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant remove product to favourites",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product removed from favourites",
		Data:       nil,
		Errors:     nil,
	})
}
