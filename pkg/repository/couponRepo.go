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

func (c *CouponDatabase) ApplayCoupon(userId, couponId int) (int, error) {
	tx := c.DB.Begin()
	//find the details corespoding to the coupon
	var coupenDetails domain.Coupons
	getcoupenDetails := `SELECT * FROM coupons WHERE id=?`
	err := tx.Raw(getcoupenDetails, couponId).Scan(&coupenDetails).Error
	if err != nil {
		tx.Rollback()
		return 0, nil
	}
	if coupenDetails.Id == 0 {
		tx.Rollback()
		return 0, fmt.Errorf("no such coupon")
	}

	//find the details of the user cart
	var cartDetails domain.Carts
	getCartDetails := `SELECT * FROM carts WHERE users_id=$1`
	err = tx.Raw(getCartDetails, userId).Scan(&cartDetails).Error
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	//check wether the coupon is valid
	if coupenDetails.ExpirationDate.Before(time.Now()) {
		tx.Rollback()
		return 0, fmt.Errorf("token expiry has end")
	}

	//chech this coupen is alredy applied by the user

	//check the tottal price meat the require ment to applay the coupen
	if coupenDetails.MinimumPurchaseAmount < float64(cartDetails.Tottal) {
		tx.Rollback()
		return 0, fmt.Errorf("need minimum %f to applay coupen", coupenDetails.MinimumPurchaseAmount)
	}
	//find the discount amount
	discountAmount := (cartDetails.Tottal / 100) * int(coupenDetails.DiscountPercent)
	//check the discount amonunt is less than the maximum discount amount
	if discountAmount > int(coupenDetails.DiscountMaximumAmount) {
		discountAmount = int(coupenDetails.DiscountMaximumAmount)
	}

	//applay the discount amount to the cart tottal
	var discountPrice int
	updateTotal := `UPDATE carts SET tottal=carts.tottal-$1 WHERE users_id=$2 RETURNING tottal`
	err = tx.Raw(updateTotal, discountAmount, userId).Scan(&discountPrice).Error
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("token expiry has end")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return discountPrice, nil

}
