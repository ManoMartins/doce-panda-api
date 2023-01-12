package coupon

import (
	"doce-panda/domain/coupon/repository"
	"doce-panda/usecase/coupon/dtos"
	"fmt"
)

type UpdateStatusCouponUseCase struct {
	couponRepository repository.CouponRepositoryInterface
}

func NewUpdateStatusCouponUseCase(couponRepository repository.CouponRepositoryInterface) *UpdateStatusCouponUseCase {
	return &UpdateStatusCouponUseCase{couponRepository: couponRepository}
}

func (c UpdateStatusCouponUseCase) Execute(input dtos.InputUpdateStatusCouponDto) error {
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
