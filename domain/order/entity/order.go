package entity

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Order struct {
	ID           string      `json:"id" validate:"required"`
	Items        []OrderItem `json:"items" validate:"required"`
	TotalInCents int         `json:"totalInCents"`
}

type OrderInterface interface {
	Validate(props Order) error
}

func NewOrder(order Order) (*Order, error) {
	o := Order{
		ID:           order.ID,
		Items:        order.Items,
		TotalInCents: order.TotalInCents,
	}

	if order.ID == "" {
		o.ID = uuid.NewString()
	}

	err := o.Validate(o)

	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (o *Order) Validate(props Order) error {
	var Validator = validator.New()

	return Validator.Struct(props)
}
