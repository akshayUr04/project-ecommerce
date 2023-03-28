package interfaces

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
)

type CouponUsecase interface {
	AddCoupon(coupon helperStruct.Coupons) error
	UpdateCoupon(coupon helperStruct.Coupons, couponId int) (domain.Coupons, error)
	DeleteCoupon(couponId int) error
	ViewCoupons() ([]domain.Coupons, error)
	ViewCoupon(couponId int) (domain.Coupons, error)
	ApplayCoupon(userId int, couponCode string) (int, error)
	RemoveCoupon(userId int) error
}
