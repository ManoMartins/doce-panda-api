package order

import (
	"doce-panda/domain/order/repository"
	"doce-panda/usecase/order/dtos"
)

type FindAllOrderByUserIdUseCase struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewFindAllOrderByUserIdUseCase(orderRepository repository.OrderRepositoryInterface) *FindAllOrderByUserIdUseCase {
	return &FindAllOrderByUserIdUseCase{orderRepository: orderRepository}
}

func (c FindAllOrderByUserIdUseCase) Execute(input dtos.InputFindAllOrderByUserIdDto) (*[]dtos.OutputFindAllOrderByUserIdDto, error) {
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
