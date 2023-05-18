package entity

import (
	"doce-panda/domain/product/entity"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type OrderItem struct {
	ID           string         `json:"id" validate:"required"`
	ProductID    string         `json:"productId" validate:"required"`
	OrderID      string         `json:"orderId"`
	Quantity     int            `json:"quantity" validate:"required,gte=0"`
	TotalInCents int            `json:"totalInCents" validate:"required,gte=0"`
	Product      entity.Product `json:"product"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}

type OrderItemInterface interface {
	Validate(props OrderItem) error
}

func NewOrderItem(orderItem OrderItem) (*OrderItem, error) {
	o := OrderItem{
		ID:           orderItem.ID,
		ProductID:    orderItem.ProductID,
		OrderID:      orderItem.OrderID,
		Quantity:     orderItem.Quantity,
		TotalInCents: orderItem.TotalInCents,
		Product:      orderItem.Product,
		CreatedAt:    orderItem.CreatedAt,
		UpdatedAt:    orderItem.UpdatedAt,
	}

	if orderItem.ID == "" {
		o.ID = uuid.NewString()
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
