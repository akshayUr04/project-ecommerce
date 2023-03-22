package repository

import (
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
