package dtos

import (
	"doce-panda/domain/product/entity"
	"time"
)

type InputCreateProductDto struct {
	Name         string            `json:"name"`
	PriceInCents int               `json:"priceInCents"`
	Description  string            `json:"description"`
	Status       entity.StatusEnum `json:"status"`
	Flavor       string            `json:"flavor"`
	Quantity     int               `json:"quantity"`
	CategoryID   string            `json:"categoryId"`
}

type OutputCreateProductDto struct {
	ID           string            `json:"ID"`
	Name         string            `json:"name"`
	PriceInCents int               `json:"priceInCents"`
	Description  string            `json:"description"`
	Status       entity.StatusEnum `json:"status"`
	Flavor       string            `json:"flavor"`
	Quantity     int               `json:"quantity"`
	Category     entity.Category   `json:"category"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
}
