package upload

import (
	"doce-panda/domain/product/entity"
	"mime/multipart"
	"time"
)

type InputUploadProductDto struct {
	ID   string                `json:"ID"`
	File *multipart.FileHeader `json:"file"`
}

type OutputUploadProductDto struct {
	ID           string            `json:"ID"`
	Name         string            `json:"name"`
	PriceInCents int               `json:"priceInCents"`
	Description  string            `json:"description"`
	Status       entity.StatusEnum `json:"status"`
	Flavor       string            `json:"flavor"`
	Quantity     int               `json:"quantity"`
	ImageUrl     string            `json:"image_url"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
}
