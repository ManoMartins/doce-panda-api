package repository

import (
	"doce-panda/domain/order/entity"
)

type OrderPendingCouponRepositoryInterface interface {
	Create(orderPendingCoupon entity.OrderPendingCoupon) error
	DeleteByOrderID(OrderID string) error
	FindByOrderId(OrderID string) (*entity.OrderPendingCoupon, error)
}
