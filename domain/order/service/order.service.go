package service

import "doce-panda/domain/order/entity"

func Total(orders []entity.OrderItem) int {
	totalInCents := 0

	for _, order := range orders {
		totalInCents += order.TotalInCents
	}

	return totalInCents
}
