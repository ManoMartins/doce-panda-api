package coupon

import (
	"doce-panda/domain/coupon/repository"
	"doce-panda/usecase/coupon/dtos"
)

type FindAllCouponUseCase struct {
	couponRepository repository.CouponRepositoryInterface
}

func NewFindAllCouponUseCase(couponRepository repository.CouponRepositoryInterface) *FindAllCouponUseCase {
	return &FindAllCouponUseCase{couponRepository: couponRepository}
}

func (c FindAllCouponUseCase) Execute() (*[]dtos.OutputFindAllCouponDto, error) {
	coupons, err := c.couponRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var output []dtos.OutputFindAllCouponDto
	for _, coupon := range *coupons {
		output = append(output, dtos.OutputFindAllCouponDto{
			ID:          coupon.ID,
			Status:      coupon.Status,
			VoucherCode: coupon.VoucherCode,
			UserID:      coupon.UserID,
			Amount:      coupon.Amount,
			CreatedAt:   coupon.CreatedAt,
			UpdatedAt:   coupon.UpdatedAt,
		})
	}

	return &output, nil
}
