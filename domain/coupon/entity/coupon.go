package entity

import (
	"doce-panda/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type StatusEnum string

const (
	NEW                 StatusEnum = "NEW"
	USED                StatusEnum = "USED"
	AWAITING_PERMISSION StatusEnum = "AWAITING_PERMISSION"
)

type Coupon struct {
	ID          string     `json:"id" validate:"required"`
	VoucherCode string     `json:"voucherCode" validate:"required"`
	Status      StatusEnum `json:"status" validate:"required,oneof='NEW' 'USED' 'AWAITING_PERMISSION'"`
	UserID      string     `json:"userId"`
	Amount      int        `json:"amount" validate:"required"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type CouponUpdateProps struct {
	VoucherCode string `json:"voucherCode" validate:"required"`
}

type CouponInterface interface {
	Validate() error
	Update(couponUpdate CouponUpdateProps) error
	UseCoupon() error
	AcceptToUseCoupon() error
}

func NewCoupon(coupon Coupon) (*Coupon, error) {
	c := Coupon{
		ID:          coupon.ID,
		VoucherCode: coupon.VoucherCode,
		Status:      coupon.Status,
		Amount:      coupon.Amount,
		UserID:      coupon.UserID,
		CreatedAt:   coupon.CreatedAt,
		UpdatedAt:   coupon.UpdatedAt,
	}

	if coupon.ID == "" {
		c.ID = uuid.NewString()
	}

	if coupon.VoucherCode == "" {
		c.VoucherCode = utils.RandStringRunes(5) + fmt.Sprintf("%v", coupon.Amount)
	}

	if coupon.Status == "" {
		c.Status = AWAITING_PERMISSION
	}

	err := c.Validate()

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Coupon) Validate() error {
	var Validator = validator.New()

	return Validator.Struct(c)
}

func (c *Coupon) Update(couponUpdate CouponUpdateProps) error {
	c.VoucherCode = couponUpdate.VoucherCode

	err := c.Validate()

	if err != nil {
		return err
	}

	return nil
}

func (c *Coupon) UseCoupon() error {
	if c.Status == USED {
		return fmt.Errorf("O cupom já foi utilizado")
	}

	if c.Status == AWAITING_PERMISSION {
		return fmt.Errorf("O cupom está sendo analisado")
	}

	c.Status = USED

	return nil
}

func (c *Coupon) AcceptToUseCoupon() error {
	if c.Status == USED {
		return fmt.Errorf("O cupom já foi utilizado")
	}

	if c.Status != AWAITING_PERMISSION {
		return fmt.Errorf("O cupom já foi analisado")
	}

	c.Status = NEW

	return nil
}
