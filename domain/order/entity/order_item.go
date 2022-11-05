package entity

import (
	"github.com/go-playground/validator/v10"
)

type OrderItem struct {
	ProductID    string `json:"productId" validate:"required"`
	Quantity     int    `json:"quantity" validate:"required,gte=0"`
	TotalInCents int    `json:"totalInCents" validate:"required,gte=0"`
}

type OrderItemInterface interface {
	Validate(props OrderItem) error
}

func NewOrderItem(orderItem OrderItem) (*OrderItem, error) {
	o := OrderItem{
		ProductID:    orderItem.ProductID,
		Quantity:     orderItem.Quantity,
		TotalInCents: orderItem.TotalInCents,
	}

	err := o.Validate(o)

	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (o *OrderItem) Validate(props OrderItem) error {
	var Validator = validator.New()

	return Validator.Struct(props)
}
