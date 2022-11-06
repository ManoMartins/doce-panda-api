package dtos

import "doce-panda/domain/order/entity"

type InputFindByIdOrderDto struct {
	ID string `json:"id"`
}

type OutputFindByIdOrderDto struct {
	ID           string             `json:"id"`
	OrderItems   []entity.OrderItem `json:"orderItems"`
	TotalInCents int                `json:"totalInCents"`
	Status       entity.StatusEnum  `json:"status"`
}
