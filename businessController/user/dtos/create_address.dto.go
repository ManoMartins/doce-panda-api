package dtos

import "time"

type InputCreateAddressDto struct {
	City         string `json:"city"`
	State        string `json:"state"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	ZipCode      string `json:"zipCode"`
	Neighborhood string `json:"neighborhood"`
	IsMain       bool   `json:"isMain"`
	UserID       string `json:"userId"`
}

type OutputCreateAddressDto struct {
	ID           string    `json:"id"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Street       string    `json:"street"`
	Number       string    `json:"number"`
	ZipCode      string    `json:"zipCode"`
	Neighborhood string    `json:"neighborhood"`
	IsMain       bool      `json:"isMain"`
	UserID       string    `json:"userId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
