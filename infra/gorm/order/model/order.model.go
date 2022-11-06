package model

import "doce-panda/domain/order/entity"

type Order struct {
	ID           string            `json:"id" gorm:"type:uuid;primary_key"`
	TotalInCents int               `json:"totalInCents" gorm:"type:integer"`
	OrderItems   []OrderItem       `json:"orderItems"`
	Status       entity.StatusEnum `json:"status" gorm:"type:varchar(255)"`
}
