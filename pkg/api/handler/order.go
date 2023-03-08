package handler

import (
	"net/http"

	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase  services.OrderUseCase
	findIdUseCase services.FindIdUseCase
}

func NewOrderHandler(orderUseCase services.OrderUseCase, findIdUseCase services.FindIdUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase:  orderUseCase,
		findIdUseCase: findIdUseCase,
	}
}

func (cr *OrderHandler) PlaceOrder(c *gin.Context) {
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
	cr.orderUseCase.PlaceOrder(userId)
}
