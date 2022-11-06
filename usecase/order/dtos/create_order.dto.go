package dtos

import "doce-panda/domain/order/entity"

type InputCreateOrderDto struct {
	OrderItems []struct {
		Quantity  int    `json:"quantity"`
		ProductID string `json:"productId"`
	} `json:"orderItems"`
}

type OutputCreateOrderDto struct {
	ID           string             `json:"id"`
	OrderItems   []entity.OrderItem `json:"orderItems"`
	TotalInCents int                `json:"totalInCents"`
	Status       entity.StatusEnum  `json:"status"`
}
