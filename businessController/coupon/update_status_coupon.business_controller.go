package coupon

import (
	"doce-panda/businessController/coupon/dtos"
	"doce-panda/domain/coupon/repository"
	"fmt"
)

type UpdateStatusCouponBusinessController struct {
	couponRepository repository.CouponRepositoryInterface
}

func NewUpdateStatusCouponBusinessController(couponRepository repository.CouponRepositoryInterface) *UpdateStatusCouponBusinessController {
	return &UpdateStatusCouponBusinessController{couponRepository: couponRepository}
}

func (c UpdateStatusCouponBusinessController) Execute(input dtos.InputUpdateStatusCouponDto) error {
	coupon, err := c.couponRepository.FindById(input.ID)

	if err != nil {
		return err
	}

	if coupon.UserID != input.UserID {
		return fmt.Errorf("O cupom não pertence ao usuário")
	}

	err = coupon.UseCoupon()

	if err != nil {
		return err
	}

	err = c.couponRepository.UpdateStatus(*coupon)

	if err != nil {
		return err
	}

	return nil
}
