package handler

import (
	"net/http"
	"strconv"

	"github.com/akshayur04/project-ecommerce/pkg/api/handlerUtil"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type FavouriteHandler struct {
	favouritesUsecase services.FavouritesUsecase
}

func NewFavouritesHandler(favouritesUsecase services.FavouritesUsecase) *FavouriteHandler {
	return &FavouriteHandler{
		favouritesUsecase: favouritesUsecase,
	}
}

// AddToFavourites
// @Summary User can add product item to favourites
// @ID add-to-favourites
// @Description User can add product item to favourites
// @Tags Favourites
// @Accept json
// @Produce json
// @Param productId path string true "ID of the product item to be added to wishlist"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/favourites/add/{productId} [post]
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
	userId, err := handlerUtil.GetUserIdFromContext(c)
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

// ReomveFavourites
// @Summary User can remove product item from favourites
// @ID remove-from-favourites
// @Description User can remove product item from favourites
// @Tags Favourites
// @Accept json
// @Produce json
// @Param productId path string true "ID of the product item to be added to wishlist"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/favourites/remove/{productId} [delete]
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
	userId, err := handlerUtil.GetUserIdFromContext(c)
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

// ViewFavourites
// @Summary User can view items in favourites
// @ID view-favourites
// @Description User view product items in favourites
// @Tags Favourites
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/favourites/view/ [get]
func (cr *FavouriteHandler) ViewFavourites(c *gin.Context) {
	userId, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	favourites, err := cr.favouritesUsecase.ViewFavourites(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant view favourites",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "products in favourites are",
		Data:       favourites,
		Errors:     nil,
	})
}
