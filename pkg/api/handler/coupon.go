package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/akshayur04/project-ecommerce/pkg/api/handlerUtil"
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	couponusecase services.CouponUsecase
}

func NewCouponHandler(couponusecase services.CouponUsecase) *CouponHandler {
	return &CouponHandler{
		couponusecase: couponusecase,
	}
}

// CreateCoupon
// @Summary Admin can create new coupon
// @ID create-coupon
// @Description Admin can create new coupons
// @Tags Coupon
// @Accept json
// @Produce json
// @Param new_coupon_details body helperStruct.Coupons true "details of new coupon to be created"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/coupon/create [post]
func (cr *CouponHandler) CreateCoupon(c *gin.Context) {
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

// UpdateCoupon
// @Summary Admin can update existing coupon
// @ID update-coupon
// @Description Admin can update existing coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param couponId path string true "Coupon ID"
// @Param coupon_details body helperStruct.Coupons true "details of coupon to be updated"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/coupon/update/{couponId} [patch]
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

// DeleteCoupon
// @Summary Admin can delete existing coupon
// @ID delete-coupon
// @Description Admin can delete existing coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param coupon_id path string true "details of coupon to be updated"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/coupon/delete/{couponId} [delete]
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

// ViewAllCoupons
// @Summary Admins and users can see all available coupons
// @ID view-coupons
// @Description Admins and users can see all available coupons
// @Tags Coupon
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/coupon/viewall  [get]
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

// ViewCouponByID
// @Summary Admins and users can see coupon with coupon id
// @ID view-coupon-by-id
// @Description Admins and users can see coupon with id
// @Tags Coupon
// @Accept json
// @Produce json
// @Param couponId path string true "coupon_id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/coupon/view/{couponId} [get]
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

// ApplayCoupon
// @Summary User can add a coupon to the cart
// @ID applay-coupon-to-cart
// @Description User can add coupon to the cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param coupon_id path string true "coupon_id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/coupon/applay/{coupon_id} [patch]
func (cr *CouponHandler) ApplayCoupon(c *gin.Context) {
	userId, err := handlerUtil.GetUserIdFromContext(c)
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

// ReomoveCoupon
// @Summary User can remove the coupon that added to the cart
// @ID remove-coupon-to-cart
// @Description User can add coupon to the cart
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/coupon/remove [patch]
func (cr *CouponHandler) RemoveCoupon(c *gin.Context) {
	userId, err := handlerUtil.GetUserIdFromContext(c)
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
