package dtos

import (
	"doce-panda/domain/coupon/entity"
	"time"
)

type InputFindByVoucherCodeCouponDto struct {
	VoucherCode string `json:"voucherCode"`
}

type OutputFindByVoucherCodeCouponDto struct {
	ID          string            `json:"id"`
	Status      entity.StatusEnum `json:"status"`
	VoucherCode string            `json:"voucherCode"`
	UserID      string            `json:"userId"`
	Amount      int               `json:"amount"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}
