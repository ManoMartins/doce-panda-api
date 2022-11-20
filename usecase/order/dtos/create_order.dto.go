package dtos

import "doce-panda/domain/order/entity"

type InputCreateOrderDto struct {
	AddressID  string `json:"addressId"`
	UserID     string `json:"userId"`
	OrderItems []struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	} `json:"orderItems"`
	Payment []struct {
		PaymentID    string `json:"paymentId"`
		TotalInCents int    `json:"totalInCents"`
	} `json:"payment"`
}

type OutputCreateOrderDto struct {
	ID           string             `json:"id"`
	OrderItems   []entity.OrderItem `json:"orderItems"`
	TotalInCents int                `json:"totalInCents"`
	Status       entity.StatusEnum  `json:"status"`
}
