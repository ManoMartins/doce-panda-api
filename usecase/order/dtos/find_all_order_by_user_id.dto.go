package dtos

import (
	"doce-panda/domain/order/entity"
	"time"
)

type InputFindAllOrderByUserIdDto struct {
	UserID string `json:"userId"`
}

type OutputFindAllOrderByUserIdDto struct {
	ID           string             `json:"id"`
	OrderItems   []entity.OrderItem `json:"orderItems"`
	TotalInCents int                `json:"totalInCents"`
	Status       entity.StatusEnum  `json:"status"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
}
