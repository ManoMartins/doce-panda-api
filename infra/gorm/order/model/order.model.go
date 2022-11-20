package model

import (
	"doce-panda/domain/order/entity"
	addressEntity "doce-panda/domain/user/entity"
	paymentModel "doce-panda/infra/gorm/payment/model"
	"time"
)

type Order struct {
	ID           string                    `json:"id" gorm:"type:uuid;primary_key"`
	TotalInCents int                       `json:"totalInCents" gorm:"type:integer"`
	OrderItems   []OrderItem               `json:"orderItems"`
	Status       entity.StatusEnum         `json:"status" gorm:"type:varchar(255)"`
	Payments     []paymentModel.CreditCard `json:"payments" gorm:"many2many:order_payments"`
	Address      addressEntity.Address     `json:"address" gorm:"ForeignKey:AddressID"`
	AddressID    string                    `json:"addressId" gorm:"column:address_id;type:uuid;notnull"`
	User         addressEntity.User        `json:"user" gorm:"ForeignKey:UserID"`
	UserID       string                    `json:"userId" gorm:"column:user_id;type:uuid;notnull"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type OrderPayment struct {
	OrderID      string
	CreditCardID string
	TotalInCents int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
