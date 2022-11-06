package order

import (
	"doce-panda/domain/order/repository"
	"doce-panda/usecase/order/dtos"
)

type FindAllOrderUseCase struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewFindAllOrderUseCase(orderRepository repository.OrderRepositoryInterface) *FindAllOrderUseCase {
	return &FindAllOrderUseCase{orderRepository: orderRepository}
}

func (c FindAllOrderUseCase) Execute() (*[]dtos.OutputFindAllOrderDto, error) {
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
		})
	}

	return &output, nil
}
