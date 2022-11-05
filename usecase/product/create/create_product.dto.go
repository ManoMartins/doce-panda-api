package create

import "time"

type StatusEnum string

type InputCreateProductDto struct {
	Name         string     `json:"name"`
	PriceInCents int        `json:"priceInCents"`
	Description  string     `json:"description"`
	Status       StatusEnum `json:"status"`
	Flavor       string     `json:"flavor"`
	Quantity     int        `json:"quantity"`
}

type OutputCreateProductDto struct {
	ID           string     `json:"ID"`
	Name         string     `json:"name"`
	PriceInCents int        `json:"priceInCents"`
	Description  string     `json:"description"`
	Status       StatusEnum `json:"status"`
	Flavor       string     `json:"flavor"`
	Quantity     int        `json:"quantity"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}
