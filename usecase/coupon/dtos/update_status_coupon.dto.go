package dtos

import (
	"doce-panda/domain/coupon/entity"
)

type InputUpdateStatusCouponDto struct {
	ID     string            `json:"id"`
	UserID string            `json:"userId"`
	Status entity.StatusEnum `json:"status"`
}
