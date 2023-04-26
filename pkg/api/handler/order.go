package handler

import (
	"net/http"
	"strconv"

	"github.com/akshayur04/project-ecommerce/pkg/api/handlerUtil"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(orderUseCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}

// OrderAll
// @Summary Buy all items from the user's cart
// @ID buyAll
// @Description This endpoint allows a user to purchase all items in their cart
// @Tags Order
// @Accept json
// @Produce json
// @Param payment_id path string true "payment_id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/order/orderall/{payment_id} [post]
func (cr *OrderHandler) OrderAll(c *gin.Context) {
	paramsId := c.Param("payment_id")
	paymentTypeId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
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
	order, err := cr.orderUseCase.OrderAll(userId, paymentTypeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant place order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "orderplaced",
		Data:       order,
		Errors:     nil,
	})
}

// CancelOrder
// @Summary Cancels a specific order for the currently logged in user
// @ID cancel-order
// @Description Endpoint for cancelling an order associated with a user
// @Tags Order
// @Accept json
// @Produce json
// @Param orderId path int true "ID of the order to be cancelled"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/order/cancel/{orderId} [patch]
func (cr *OrderHandler) UserCancelOrder(c *gin.Context) {
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
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.orderUseCase.UserCancelOrder(orderId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't cancel order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order canceld",
		Data:       nil,
		Errors:     nil,
	})
}

// ViewOrderByID function retrieves order details for a given order ID, if authorized.
// @Summary Retrieves order details for a given order ID, if authorized.
// @ID view-order-by-id
// @Description This function handles requests for retrieving the details of a specific order identified by its order ID.
// @Tags Order
// @Accept json
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} response.Response "Successfully fetched order details"
// @Failure 400 {object} response.Response "Failed to fetch order details"
// @Router /user/order/view/{orderId} [get]
func (cr *OrderHandler) ListOrder(c *gin.Context) {
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
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	order, err := cr.orderUseCase.ListOrder(userId, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order ",
		Data:       order,
		Errors:     nil,
	})
}

// ViewAllOrders
// @Summary Retrieves all orders of currently logged in user
// @ID view-all-orders
// @Description Endpoint for getting all orders associated with a user
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/order/listall [get]
func (cr *OrderHandler) ListAllOrders(c *gin.Context) {
	Id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	orders, err := cr.orderUseCase.ListAllOrders(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order ",
		Data:       orders,
		Errors:     nil,
	})
}

// ReturnOrder
// @Summary Return a specific order for the currently logged in user
// @ID return-order
// @Description Endpoint for Returning an order associated with a user
// @Tags Order
// @Accept json
// @Produce json
// @Param orderId path int true "ID of the order to be cancelled"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/order/return/{orderId} [patch]
func (cr *OrderHandler) ReturnOrder(c *gin.Context) {
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
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	returnAmount, err := cr.orderUseCase.ReturnOrder(userId, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't return order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order returnd ",
		Data:       returnAmount,
		Errors:     nil,
	})
}

func (cr *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("order_id")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.orderUseCase.UpdateOrder(orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't update order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order updated ",
		Data:       nil,
		Errors:     nil,
	})
}
