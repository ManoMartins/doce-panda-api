package find

import "time"

type InputFindProductDto struct {
	ID string
}

type StatusEnum string

type OutputFindProductDto struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	PriceInCents int        `json:"priceInCents"`
	Status       StatusEnum `json:"status"`
	Description  string     `json:"description"`
	Flavor       string     `json:"flavor"`
	Quantity     int        `json:"quantity"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
