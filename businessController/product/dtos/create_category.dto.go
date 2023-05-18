package dtos

import "time"

type InputCreateCategoryDto struct {
	Description string `json:"description"`
}

type OutputCreateCategoryDto struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
