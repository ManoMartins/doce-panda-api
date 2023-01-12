package coupon

import (
	"doce-panda/domain/coupon/repository"
	"doce-panda/usecase/coupon/dtos"
)

type FindByIdCouponUseCase struct {
	couponRepository repository.CouponRepositoryInterface
}

func NewFindByIdCouponUseCase(couponRepository repository.CouponRepositoryInterface) *FindByIdCouponUseCase {
	return &FindByIdCouponUseCase{couponRepository: couponRepository}
}

func (c FindByIdCouponUseCase) Execute(input dtos.InputFindByIdCouponDto) (*dtos.OutputFindByIdCouponDto, error) {
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
