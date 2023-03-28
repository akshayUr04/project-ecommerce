package repository

import (
	"fmt"
	"time"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type CouponDatabase struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &CouponDatabase{DB}
}

func (c *CouponDatabase) AddCoupon(coupon helperStruct.Coupons) error {
	createCoupen := `INSERT INTO coupons (code, discount_percent,discount_maximum_amount,minimum_purchase_amount,expiration_date)
		VALUES($1,$2,$3,$4,$5)`
	err := c.DB.Exec(createCoupen,
		coupon.Code,
		coupon.DiscountPercent,
		coupon.DiscountMaximumAmount,
		coupon.MinimumPurchaseAmount,
		coupon.ExpirationDate).Error
	return err
}

func (c *CouponDatabase) UpdateCoupon(coupon helperStruct.Coupons, couponId int) (domain.Coupons, error) {
	var updatedCoupen domain.Coupons
	updateCoupon := `UPDATE coupons SET code=$1, discount_percent=$2,discount_maximum_amount=$3,minimum_purchase_amount=$4,expiration_date=$5
		 WHERE id=$6
		 RETURNING *`
	err := c.DB.Raw(updateCoupon,
		coupon.Code,
		coupon.DiscountPercent,
		coupon.DiscountMaximumAmount,
		coupon.MinimumPurchaseAmount,
		coupon.ExpirationDate,
		couponId).
		Scan(&updatedCoupen).
		Error
	return updatedCoupen, err
}

func (c *CouponDatabase) DeleteCoupon(couponId int) error {
	deleteCoupon := `DELETE FROM coupons WHERE id=?`
	err := c.DB.Exec(deleteCoupon, couponId).Error
	return err
}

func (c *CouponDatabase) ViewCoupons() ([]domain.Coupons, error) {
	var coupens []domain.Coupons
	fetchDetails := `SELECT * FROM coupons`
	err := c.DB.Raw(fetchDetails).Scan(&coupens).Error
	return coupens, err

}

func (c *CouponDatabase) ViewCoupon(couponId int) (domain.Coupons, error) {
	var coupon domain.Coupons
	fetchCoupenDetails := `SELECT * FORM coupons WHERE id=$1`
	err := c.DB.Raw(fetchCoupenDetails, couponId).Scan(&coupon).Error
	return coupon, err

}

func (c *CouponDatabase) ApplayCoupon(userId int, couponCode string) (int, error) {
	tx := c.DB.Begin()
	// find the coupon details
	var couponDetails domain.Coupons
	findCouponDetails := `SELECT * FROM coupons WHERE code=$1`
	err := tx.Raw(findCouponDetails, couponCode).Scan(&couponDetails).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if couponDetails.ExpirationDate.Before(time.Now()) {
		tx.Rollback()
		return 0, fmt.Errorf("coupon expired")
	}

	// check whether the coupen is alredy used by the user in any other previous orders
	var isUsed bool
	checkIsUsed := `SELECT EXISTS(SELECT 1 FROM orders WHERE coupon_id = $1 AND user_id = $2)`
	err = tx.Raw(checkIsUsed, couponDetails.Id, userId).Scan(&isUsed).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if isUsed {
		tx.Rollback()
		return 0, fmt.Errorf("coupen is alredy used")
	}
	// check whether the coupen is alresy added to the cart
	var cartDetails domain.Carts
	getCartDetails := `SELECT * FROM carts WHERE user_id=?`
	err = tx.Raw(getCartDetails, userId).Scan(&cartDetails).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if cartDetails.CouponId == couponDetails.Id {
		tx.Rollback()
		return 0, fmt.Errorf("coupen is applied to cart")
	}

	//check there is some thing inside the cart
	if cartDetails.SubTotal == 0 {
		tx.Rollback()
		return 0, fmt.Errorf("no product is in the cart to applay coupen")
	}

	// check whether the coupen minimum purchase value is greater than or equal to the order amount
	if cartDetails.SubTotal <= int(couponDetails.MinimumPurchaseAmount) {
		tx.Rollback()
		return 0, fmt.Errorf("need minimum %v in cart", couponDetails.MinimumPurchaseAmount)
	}

	//check the discount amonunt is less than the maximum discount amount
	discountAmount := (cartDetails.SubTotal / 100) * int(couponDetails.DiscountPercent)
	if discountAmount > int(couponDetails.DiscountMaximumAmount) {
		discountAmount = int(couponDetails.DiscountMaximumAmount)
	}

	// update the cart total with the subtotal - discount amount if the cart alredy have anything in cart
	updateCart := `UPDATE carts SET total=$1,coupon_id=$2 WHERE id=$3`
	err = tx.Exec(updateCart, cartDetails.SubTotal-discountAmount, couponDetails.Id, cartDetails.Id).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return cartDetails.Total, nil

}

func (c *CouponDatabase) RemoveCoupon(userId int) error {
	tx := c.DB.Begin()
	//get the details of the cart
	var cartDetails domain.Carts
	getCartDetails := `SELECT * FROM carts WHERE user_id=$1`
	err := tx.Raw(getCartDetails, userId).Scan(&cartDetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//find any coupon is added
	if cartDetails.CouponId == 0 {
		tx.Rollback()
		return fmt.Errorf("no coupon to remove")
	}
	//if added remove the coupon
	removeCoupon := `UPDATE carts SET coupon_id=0, total = sub_total WHERE user_id=$1`
	err = tx.Exec(removeCoupon, userId).Error
	if cartDetails.CouponId == 0 {
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
