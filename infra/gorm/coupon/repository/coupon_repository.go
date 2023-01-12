package repository

import (
	"doce-panda/domain/coupon/entity"
	"doce-panda/infra/gorm/coupon/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type CouponRepositoryDb struct {
	Db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) *CouponRepositoryDb {
	return &CouponRepositoryDb{Db: db}
}

func (c CouponRepositoryDb) FindById(ID string) (*entity.Coupon, error) {
	var couponModel model.Coupon

	c.Db.First(&couponModel, "id = ?", ID)

	if couponModel.ID == "" {
		return nil, fmt.Errorf("O cupom não foi encontrado")
	}

	return entity.NewCoupon(entity.Coupon{
		ID:          couponModel.ID,
		VoucherCode: couponModel.VoucherCode,
		Status:      couponModel.Status,
		UserID:      couponModel.UserID,
		Amount:      couponModel.Amount,
		CreatedAt:   couponModel.CreatedAt,
		UpdatedAt:   couponModel.UpdatedAt,
	})
}

func (c CouponRepositoryDb) FindAll() (*[]entity.Coupon, error) {
	var couponsModel []model.Coupon

	err := c.Db.Find(&couponsModel).Error

	if err != nil {
		return nil, err
	}

	var coupons []entity.Coupon

	for _, couponModel := range couponsModel {
		coupon, err := entity.NewCoupon(entity.Coupon{
			ID:          couponModel.ID,
			VoucherCode: couponModel.VoucherCode,
			Status:      couponModel.Status,
			UserID:      couponModel.UserID,
			Amount:      couponModel.Amount,
			CreatedAt:   couponModel.CreatedAt,
			UpdatedAt:   couponModel.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		coupons = append(coupons, *coupon)
	}

	return &coupons, nil
}

func (c CouponRepositoryDb) Delete(ID string) error {
	coupon := entity.Coupon{ID: ID}

	err := c.Db.Delete(&coupon).Error

	if err != nil {
		return err
	}

	return nil
}

func (c CouponRepositoryDb) UpdateStatus(coupon entity.Coupon) error {
	couponModel := model.Coupon{
		ID:          coupon.ID,
		VoucherCode: coupon.VoucherCode,
		Status:      coupon.Status,
		UserID:      coupon.UserID,
		Amount:      coupon.Amount,
		CreatedAt:   coupon.CreatedAt,
		UpdatedAt:   coupon.UpdatedAt,
	}

	err := c.Db.Save(&couponModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (c CouponRepositoryDb) Create(coupon entity.Coupon) error {
	couponModel := model.Coupon{
		ID:          coupon.ID,
		VoucherCode: coupon.VoucherCode,
		Status:      coupon.Status,
		UserID:      coupon.UserID,
		Amount:      coupon.Amount,
		CreatedAt:   coupon.CreatedAt,
		UpdatedAt:   coupon.UpdatedAt,
	}

	err := c.Db.Create(&couponModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (c CouponRepositoryDb) FindByVoucherCode(VoucherCode string) (*entity.Coupon, error) {
	var couponModel model.Coupon

	c.Db.First(&couponModel, "voucher_code = ?", VoucherCode)

	if couponModel.ID == "" {
		return nil, fmt.Errorf("O cupom não foi encontrado")
	}

	return entity.NewCoupon(entity.Coupon{
		ID:          couponModel.ID,
		VoucherCode: couponModel.VoucherCode,
		Status:      couponModel.Status,
		UserID:      couponModel.UserID,
		Amount:      couponModel.Amount,
		CreatedAt:   couponModel.CreatedAt,
		UpdatedAt:   couponModel.UpdatedAt,
	})
}
