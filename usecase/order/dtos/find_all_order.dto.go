package dtos

import "doce-panda/domain/order/entity"

type OutputFindAllOrderDto struct {
	ID           string             `json:"id"`
	OrderItems   []entity.OrderItem `json:"orderItems"`
	TotalInCents int                `json:"totalInCents"`
	Status       entity.StatusEnum  `json:"status"`
}
