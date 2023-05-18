package coupon

import (
	"doce-panda/businessController/coupon/dtos"
	"doce-panda/domain/coupon/repository"
)

type FindAllCouponBusinessController struct {
	couponRepository repository.CouponRepositoryInterface
}

func NewFindAllCouponBusinessController(couponRepository repository.CouponRepositoryInterface) *FindAllCouponBusinessController {
	return &FindAllCouponBusinessController{couponRepository: couponRepository}
}

func (c FindAllCouponBusinessController) Execute() (*[]dtos.OutputFindAllCouponDto, error) {
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
