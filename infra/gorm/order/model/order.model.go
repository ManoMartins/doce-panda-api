package model

import (
	couponEntity "doce-panda/domain/coupon/entity"
	"doce-panda/domain/order/entity"
	addressEntity "doce-panda/domain/user/entity"
	paymentModel "doce-panda/infra/gorm/payment/model"
	"time"
)

type Order struct {
	ID            string                    `json:"id" gorm:"type:uuid;primary_key"`
	TotalInCents  int                       `json:"totalInCents" gorm:"type:integer"`
	OrderItems    []OrderItem               `json:"orderItems"`
	Status        entity.StatusEnum         `json:"status" gorm:"type:varchar(255)"`
	Payments      []paymentModel.CreditCard `json:"payments" gorm:"many2many:order_payments"`
	OrderPayments []OrderPayment            `json:"orderPayments"`
	Coupon        couponEntity.Coupon       `json:"coupon" gorm:"ForeignKey:CouponID"`
	CouponID      *string                   `json:"couponId" gorm:"column:coupon_id;type:uuid;nullable"`
	Address       addressEntity.Address     `json:"address" gorm:"ForeignKey:AddressID"`
	AddressID     string                    `json:"addressId" gorm:"column:address_id;type:uuid;notnull"`
	User          addressEntity.User        `json:"user" gorm:"ForeignKey:UserID"`
	UserID        string                    `json:"userId" gorm:"column:user_id;type:uuid;notnull"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type OrderPayment struct {
	OrderID      string
	CreditCardID string
	TotalInCents int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
