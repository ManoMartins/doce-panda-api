package model

import (
	"doce-panda/infra/gorm/coupon/model"
	"time"
)

type OrderPendingCoupon struct {
	ID        string       `json:"id" gorm:"type:uuid;primary_key"`
	Coupon    model.Coupon `json:"coupon" `
	CouponID  string       `json:"couponId" gorm:"column:coupon_id;type:uuid;notnull"`
	OrderID   string       `json:"orderId" gorm:"column:order_id;type:uuid;notnull"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
