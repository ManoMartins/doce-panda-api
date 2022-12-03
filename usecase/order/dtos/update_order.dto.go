package dtos

import "doce-panda/domain/order/entity"

type InputUpdateOrderDto struct {
	ID     string            `json:"id"`
	Status entity.StatusEnum `json:"status"`
}
