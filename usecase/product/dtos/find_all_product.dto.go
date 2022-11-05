package dtos

import (
	"doce-panda/domain/product/entity"
	"time"
)

type OutputFindAllProductDto struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	PriceInCents int               `json:"priceInCents"`
	Status       entity.StatusEnum `json:"status"`
	Description  string            `json:"description"`
	Flavor       string            `json:"flavor"`
	Quantity     int               `json:"quantity"`
	ImageUrl     string            `json:"imageUrl"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}
