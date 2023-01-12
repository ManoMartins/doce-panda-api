package dtos

import (
	"doce-panda/domain/coupon/entity"
	"time"
)

type InputFindByIdCouponDto struct {
	ID string `json:"id"`
}

type OutputFindByIdCouponDto struct {
	ID          string            `json:"id"`
	Status      entity.StatusEnum `json:"status"`
	VoucherCode string            `json:"voucherCode"`
	UserID      string            `json:"userId"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}
