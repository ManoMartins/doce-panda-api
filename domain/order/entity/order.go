package entity

import (
	"doce-panda/domain/payment/entity"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type StatusEnum string

const (
	WAITING_PAYMENT StatusEnum = "WAITING_PAYMENT"
	PREPARING       StatusEnum = "PREPARING"
	IN_TRANSIT      StatusEnum = "IN_TRANSIT"
	DELIVERED       StatusEnum = "DELIVERED"
)

type Order struct {
	ID           string              `json:"id" validate:"required"`
	OrderItems   []OrderItem         `json:"orderItems"`
	TotalInCents int                 `json:"totalInCents"`
	Status       StatusEnum          `json:"status" validate:"required,oneof='WAITING_PAYMENT' 'PREPARING' 'IN_TRANSIT' 'DELIVERED'"`
	Payments     []entity.CreditCard `json:"payments" validate:"required"`
	AddressID    string              `json:"addressId"`
	UserID       string              `json:"userId"`
}

type OrderInterface interface {
	Validate(props Order) error
	AddOrderItems(orderItems []OrderItem)
}

func NewOrder(order Order) (*Order, error) {
	o := Order{
		ID:           order.ID,
		OrderItems:   order.OrderItems,
		TotalInCents: order.TotalInCents,
		Status:       order.Status,
		Payments:     order.Payments,
		AddressID:    order.AddressID,
		UserID:       order.UserID,
	}

	if order.ID == "" {
		o.ID = uuid.NewString()
	}

	if order.Status == "" {
		o.Status = WAITING_PAYMENT
	}

	err := o.Validate()

	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (o *Order) Validate() error {
	var Validator = validator.New()

	return Validator.Struct(*o)
}

func (o *Order) AddOrderItems(orderItems []OrderItem) {
	for i := range orderItems {
		orderItems[i].OrderID = o.ID
	}

	o.OrderItems = orderItems
}
