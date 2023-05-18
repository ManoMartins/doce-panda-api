package order

import (
	"doce-panda/businessController/order/dtos"
	"doce-panda/domain/order/repository"
)

type FindAllOrderBusinessController struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewFindAllOrderBusinessController(orderRepository repository.OrderRepositoryInterface) *FindAllOrderBusinessController {
	return &FindAllOrderBusinessController{orderRepository: orderRepository}
}

func (c FindAllOrderBusinessController) Execute() (*[]dtos.OutputFindAllOrderDto, error) {
	orders, err := c.orderRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var output []dtos.OutputFindAllOrderDto
	for _, order := range *orders {
		output = append(output, dtos.OutputFindAllOrderDto{
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
