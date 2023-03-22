package usecase

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
)

type CouponUsecase struct {
	couponRepository interfaces.CouponRepository
}

func NewCouponUsecase(couponRepository interfaces.CouponRepository) services.CouponUsecase {
	return &CouponUsecase{
		couponRepository: couponRepository,
	}
}
func (c *CouponUsecase) AddCoupon(coupon helperStruct.Coupons) error {
	err := c.couponRepository.AddCoupon(coupon)
	return err
}

func (c *CouponUsecase) UpdateCoupon(coupon helperStruct.Coupons, couponId int) (domain.Coupons, error) {
	updatedCoupen, err := c.couponRepository.UpdateCoupon(coupon, couponId)
	return updatedCoupen, err
}

func (c *CouponUsecase) DeleteCoupon(couponId int) error {
	err := c.couponRepository.DeleteCoupon(couponId)
	return err
}

func (c *CouponUsecase) ApplayCoupon(userId, couponId int) (int, error) {
	discountRate, err := c.couponRepository.ApplayCoupon(userId, couponId)
	return discountRate, err
}
