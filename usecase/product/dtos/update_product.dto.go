package dtos

import "time"

type StatusEnum string

type InputUpdateProductDto struct {
	ID           string     `json:"ID"`
	Name         string     `json:"name"`
	PriceInCents int        `json:"priceInCents"`
	Status       StatusEnum `json:"status"`
	Description  string     `json:"description"`
	Flavor       string     `json:"flavor"`
	Quantity     int        `json:"quantity"`
}

type OutputUpdateProductDto struct {
	ID           string     `json:"ID"`
	Name         string     `json:"name"`
	PriceInCents int        `json:"priceInCents"`
	Description  string     `json:"description"`
	Status       StatusEnum `json:"status"`
	Flavor       string     `json:"flavor"`
	Quantity     int        `json:"quantity"`
	ImageUrl     string     `json:"ImageUrl"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}
