package entity

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type StatusEnum string

const (
	DISABLED StatusEnum = "DISABLED"
	ENABLED  StatusEnum = "ENABLED"
)

type ProductProps struct {
	Name         string
	PriceInCents int
	Description  string
	Flavor       string
	Quantity     int
}

type Product struct {
	ID           string     `json:"id" validate:"required"`
	Name         string     `json:"name" validate:"required"`
	PriceInCents int        `json:"priceInCents" validate:"required"`
	Status       StatusEnum `json:"status" validate:"required"`
	Description  string     `json:"description" validate:"required"`
	Flavor       string     `json:"flavor" validate:"required"`
	Quantity     int        `json:"quantity" validate:"required"`
	ImageUrl     string     `json:"image_url"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

type ProductInterface interface {
	Enable() error
	Disable() error
	Validate() error
}

func NewProduct(product ProductProps) (*Product, error) {
	p := &Product{
		ID:           uuid.NewString(),
		Status:       DISABLED,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		Description:  product.Description,
	}

	err := p.Validate()

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Validate() error {
	var Validator = validator.New()

	return Validator.Struct(p)
}

func (p *Product) Disable() error {
	if p.Status == DISABLED {
		return fmt.Errorf("O produto já está desativado")
	}

	p.Status = DISABLED

	return nil
}

func (p *Product) Enable() error {
	if p.Status == ENABLED {
		return fmt.Errorf("O produto já está ativado")
	}

	p.Status = ENABLED

	return nil
}
