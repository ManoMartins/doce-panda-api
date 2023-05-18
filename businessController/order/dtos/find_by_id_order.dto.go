package dtos

import (
	"doce-panda/domain/order/entity"
	userEntity "doce-panda/domain/user/entity"
	"time"
)

type InputFindByIdOrderDto struct {
	ID string `json:"id"`
}

type OutputFindByIdOrderDto struct {
	ID           string             `json:"id"`
	OrderItems   []entity.OrderItem `json:"orderItems"`
	TotalInCents int                `json:"totalInCents"`
	Status       entity.StatusEnum  `json:"status"`
	Address      userEntity.Address `json:"address"`
	User         userEntity.User    `json:"user"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
}
