package order

import (
	"doce-panda/domain/order/repository"
	"doce-panda/usecase/order/dtos"
)

type FindByIdOrderUseCase struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewFindByIdOrderUseCase(orderRepository repository.OrderRepositoryInterface) *FindByIdOrderUseCase {
	return &FindByIdOrderUseCase{orderRepository: orderRepository}
}

func (c FindByIdOrderUseCase) Execute(input dtos.InputFindByIdOrderDto) (*dtos.OutputFindByIdOrderDto, error) {
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
