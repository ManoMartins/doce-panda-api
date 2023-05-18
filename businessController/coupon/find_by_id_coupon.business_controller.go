package coupon

import (
	"doce-panda/businessController/coupon/dtos"
	"doce-panda/domain/coupon/repository"
)

type FindByIdCouponBusinessController struct {
	couponRepository repository.CouponRepositoryInterface
}

func NewFindByIdCouponBusinessController(couponRepository repository.CouponRepositoryInterface) *FindByIdCouponBusinessController {
	return &FindByIdCouponBusinessController{couponRepository: couponRepository}
}

func (c FindByIdCouponBusinessController) Execute(input dtos.InputFindByIdCouponDto) (*dtos.OutputFindByIdCouponDto, error) {
	coupon, err := c.couponRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputFindByIdCouponDto{
		ID:          coupon.ID,
		Status:      coupon.Status,
		VoucherCode: coupon.VoucherCode,
		UserID:      coupon.UserID,
		CreatedAt:   coupon.CreatedAt,
		UpdatedAt:   coupon.UpdatedAt,
	}, nil
}
