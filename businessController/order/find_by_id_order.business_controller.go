package order

import (
	"doce-panda/businessController/order/dtos"
	"doce-panda/domain/order/repository"
)

type FindByIdOrderBusinessController struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewFindByIdOrderBusinessController(orderRepository repository.OrderRepositoryInterface) *FindByIdOrderBusinessController {
	return &FindByIdOrderBusinessController{orderRepository: orderRepository}
}

func (c FindByIdOrderBusinessController) Execute(input dtos.InputFindByIdOrderDto) (*dtos.OutputFindByIdOrderDto, error) {
	order, err := c.orderRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputFindByIdOrderDto{
		ID:           order.ID,
		OrderItems:   order.OrderItems,
		TotalInCents: order.TotalInCents,
		Status:       order.Status,
		Address:      *order.Address,
		User:         *order.User,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}, nil
}
