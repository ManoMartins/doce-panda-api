package dtos

import (
	"doce-panda/domain/coupon/entity"
	"time"
)

type OutputFindAllCouponDto struct {
	ID          string            `json:"id"`
	Status      entity.StatusEnum `json:"status"`
	VoucherCode string            `json:"voucherCode"`
	UserID      string            `json:"userId"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}
