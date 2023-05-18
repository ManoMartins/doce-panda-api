package repository

import (
	"doce-panda/domain/order/entity"
	"doce-panda/infra/gorm/order/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type OrderPendingCouponDb struct {
	Db *gorm.DB
}

func NewOrderPendingCouponRepository(db *gorm.DB) *OrderPendingCouponDb {
	return &OrderPendingCouponDb{Db: db}
}

func (o OrderPendingCouponDb) Create(orderPendingCoupon entity.OrderPendingCoupon) error {
	orderPendingCouponModel := model.OrderPendingCoupon{
		ID:        orderPendingCoupon.ID,
		CouponID:  orderPendingCoupon.CouponID,
		OrderID:   orderPendingCoupon.OrderID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := o.Db.Create(&orderPendingCouponModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (o OrderPendingCouponDb) DeleteByOrderID(OrderID string) error {
	orderPendingCoupon := entity.OrderPendingCoupon{OrderID: OrderID}

	err := o.Db.Delete(&orderPendingCoupon).Error

	if err != nil {
		return err
	}

	return nil
}

func (o OrderPendingCouponDb) FindByOrderId(OrderID string) (*entity.OrderPendingCoupon, error) {
	var orderPendingCoupon model.OrderPendingCoupon

	o.Db.First(&orderPendingCoupon, "order_id = ?", OrderID)

	if orderPendingCoupon.ID == "" {
		return nil, fmt.Errorf("O cupom não foi encontrado")
	}

	return entity.NewOrderPendingCoupon(entity.OrderPendingCoupon{
		ID:        orderPendingCoupon.ID,
		OrderID:   orderPendingCoupon.OrderID,
		CouponID:  orderPendingCoupon.CouponID,
		CreatedAt: orderPendingCoupon.CreatedAt,
		UpdatedAt: orderPendingCoupon.UpdatedAt,
	})
}
