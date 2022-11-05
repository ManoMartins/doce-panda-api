package model

import (
	"doce-panda/domain/product/entity"
	"time"
)

type Product struct {
	ID           string            `json:"id" gorm:"type:uuid;primary_key"`
	Name         string            `json:"name" gorm:"type:varchar(255)"`
	PriceInCents int               `json:"priceInCents" gorm:"type:integer"`
	Status       entity.StatusEnum `json:"status" gorm:"type:varchar(255)"`
	Description  string            `json:"description" gorm:"type:varchar(255)"`
	Flavor       string            `json:"flavor" gorm:"type:varchar(255)"`
	Quantity     int               `json:"quantity" gorm:"type:integer"`
	ImageUrl     string            `json:"imageUrl"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
}
