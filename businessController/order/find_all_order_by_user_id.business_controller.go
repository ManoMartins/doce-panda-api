package order

import (
	"doce-panda/businessController/order/dtos"
	"doce-panda/domain/order/repository"
)

type FindAllOrderByUserIdBusinessController struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewFindAllOrderByUserIdBusinessController(orderRepository repository.OrderRepositoryInterface) *FindAllOrderByUserIdBusinessController {
	return &FindAllOrderByUserIdBusinessController{orderRepository: orderRepository}
}

func (c FindAllOrderByUserIdBusinessController) Execute(input dtos.InputFindAllOrderByUserIdDto) (*[]dtos.OutputFindAllOrderByUserIdDto, error) {
	orders, err := c.orderRepository.FindAllByUserId(input.UserID)

	if err != nil {
		return nil, err
	}

	var output []dtos.OutputFindAllOrderByUserIdDto
	for _, order := range *orders {
		output = append(output, dtos.OutputFindAllOrderByUserIdDto{
			ID:           order.ID,
			OrderItems:   order.OrderItems,
			TotalInCents: order.TotalInCents,
			Status:       order.Status,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
		})
	}

	return &output, nil
}
