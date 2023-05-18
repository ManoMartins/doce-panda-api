package entity

import (
	"github.com/google/uuid"
	"time"
)

type OrderPendingCoupon struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"orderId"`
	CouponID  string    `json:"couponId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type OrderPendingCouponInterface interface {
}

func NewOrderPendingCoupon(orderPendingCoupon OrderPendingCoupon) (*OrderPendingCoupon, error) {
	o := OrderPendingCoupon{
		ID:       orderPendingCoupon.ID,
		OrderID:  orderPendingCoupon.OrderID,
		CouponID: orderPendingCoupon.CouponID,
	}

	if orderPendingCoupon.ID == "" {
		o.ID = uuid.NewString()
	}

	return &o, nil
}
