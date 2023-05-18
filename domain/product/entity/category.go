package entity

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID          string
	Description string
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewCategory(category Category) (*Category, error) {
	c := Category{
		ID:          category.ID,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}

	if category.ID == "" {
		c.ID = uuid.NewString()
	}

	return &c, nil
}
