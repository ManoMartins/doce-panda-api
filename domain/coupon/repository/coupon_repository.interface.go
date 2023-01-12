package repository

import "doce-panda/domain/coupon/entity"

type CouponRepositoryInterface interface {
	FindByVoucherCode(VoucherCode string) (*entity.Coupon, error)
	FindById(ID string) (*entity.Coupon, error)
	FindAll() (*[]entity.Coupon, error)
	Delete(ID string) error
	UpdateStatus(coupon entity.Coupon) error
	Create(coupon entity.Coupon) error
}
