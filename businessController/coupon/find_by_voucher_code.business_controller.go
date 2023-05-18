package coupon

import (
	"doce-panda/businessController/coupon/dtos"
	"doce-panda/domain/coupon/repository"
)

type FindByVoucherCodeCouponBusinessController struct {
	couponRepository repository.CouponRepositoryInterface
}

func NewFindByVoucherCodeCouponBusinessController(couponRepository repository.CouponRepositoryInterface) *FindByVoucherCodeCouponBusinessController {
	return &FindByVoucherCodeCouponBusinessController{couponRepository: couponRepository}
}

func (c FindByVoucherCodeCouponBusinessController) Execute(input dtos.InputFindByVoucherCodeCouponDto) (*dtos.OutputFindByVoucherCodeCouponDto, error) {
	coupon, err := c.couponRepository.FindByVoucherCode(input.VoucherCode)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputFindByVoucherCodeCouponDto{
		ID:          coupon.ID,
		Status:      coupon.Status,
		VoucherCode: coupon.VoucherCode,
		UserID:      coupon.UserID,
		Amount:      coupon.Amount,
		CreatedAt:   coupon.CreatedAt,
		UpdatedAt:   coupon.UpdatedAt,
	}, nil
}
