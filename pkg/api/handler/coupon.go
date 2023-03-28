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

type CouponHandler struct {
	couponusecase services.CouponUsecase
	findIdUseCase services.FindIdUseCase
}

func NewCouponHandler(couponusecase services.CouponUsecase, findIdUseCase services.FindIdUseCase) *CouponHandler {
	return &CouponHandler{
		couponusecase: couponusecase,
		findIdUseCase: findIdUseCase,
	}
}

func (cr *CouponHandler) AddCoupon(c *gin.Context) {
	var newCoupon helperStruct.Coupons
	err := c.Bind(&newCoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.couponusecase.AddCoupon(newCoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "camt't create coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupen created",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *CouponHandler) UpdateCoupon(c *gin.Context) {
	id := c.Param("couponId")
	coupenId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var newCoupon helperStruct.Coupons
	err = c.Bind(&newCoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatedCoupen, err := cr.couponusecase.UpdateCoupon(newCoupon, coupenId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't create coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupen updated",
		Data:       updatedCoupen,
		Errors:     nil,
	})
}

func (cr *CouponHandler) DeleteCoupon(c *gin.Context) {
	id := c.Param("couponId")
	couponId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.couponusecase.DeleteCoupon(couponId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't delete coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupon deleted",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *CouponHandler) ViewCoupons(c *gin.Context) {
	coupons, err := cr.couponusecase.ViewCoupons()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't finds coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupons are ",
		Data:       coupons,
		Errors:     nil,
	})
}

func (cr *CouponHandler) ViewCoupon(c *gin.Context) {
	id := c.Param("couponId")
	couponId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find couponid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	coupon, err := cr.couponusecase.ViewCoupon(couponId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't finds coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupons are ",
		Data:       coupon,
		Errors:     nil,
	})
}

func (cr *CouponHandler) ApplyCoupon(c *gin.Context) {
	cookie, err := c.Cookie("UserAuth")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userId, err := cr.findIdUseCase.FindId(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	couponCode := c.Query("c_id")
	fmt.Println(couponCode)
	discountRate, err := cr.couponusecase.ApplayCoupon(userId, couponCode)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "can't applay coupen",
				Data:       nil,
				Errors:     err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "coupen applayed",
		Data:       []interface{}{"rate after coupen applaid is ", discountRate},
		Errors:     nil,
	})

}

func (cr *CouponHandler) RemoveCoupon(c *gin.Context) {
	cookie, err := c.Cookie("UserAuth")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userId, err := cr.findIdUseCase.FindId(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.couponusecase.RemoveCoupon(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't remove coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Coupon removed",
		Data:       nil,
		Errors:     nil,
	})
}
