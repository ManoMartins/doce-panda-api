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

type Product struct {
	ID           string     `json:"id" validate:"required"`
	Name         string     `json:"name" validate:"required"`
	PriceInCents int        `json:"priceInCents" validate:"required"`
	Status       StatusEnum `json:"status" validate:"required"`
	Description  string     `json:"description" validate:"required"`
	Flavor       string     `json:"flavor" validate:"required"`
	Quantity     int        `json:"quantity" validate:"required,min=0"`
	ImageUrl     string     `json:"imageUrl"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

type ProductUpdateProps struct {
	Name         string
	PriceInCents int
	Status       StatusEnum
	Description  string
	Flavor       string
	Quantity     int
	ImageUrl     string
}

type ProductInterface interface {
	Enable() error
	Disable() error
	AddImageUrl(imageUrl string)
	Validate(props Product) error
	RemoveQuantity(amount int) error
	Update(productUpdate ProductUpdateProps) error
}

func NewProduct(product Product) (*Product, error) {
	p := Product{
		ID:           product.ID,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Status:       product.Status,
		Description:  product.Description,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		ImageUrl:     product.ImageUrl,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}

	if product.ID == "" {
		p.ID = uuid.NewString()
	}

	if product.Status == "" {
		p.Status = DISABLED
	}

	err := p.Validate(p)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Product) Validate(props Product) error {
	var Validator = validator.New()

	return Validator.Struct(props)
}

func (p *Product) Update(productUpdate ProductUpdateProps) error {
	p.Name = productUpdate.Name
	p.PriceInCents = productUpdate.PriceInCents
	p.Status = productUpdate.Status
	p.Description = productUpdate.Description
	p.Flavor = productUpdate.Flavor
	p.Quantity = productUpdate.Quantity
	p.ImageUrl = productUpdate.ImageUrl

	err := p.Validate(*p)

	if err != nil {
		return err
	}

	return nil

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

func (p *Product) AddImageUrl(imageUrl string) {
	p.ImageUrl = imageUrl
}

func (p *Product) RemoveQuantity(amount int) error {
	if p.Quantity < amount {
		return fmt.Errorf("O produto não tem quantidade suficiente")
	}

	p.Quantity -= amount

	return nil
}
