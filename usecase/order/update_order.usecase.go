package order

import (
	"doce-panda/domain/order/repository"
	"doce-panda/usecase/order/dtos"
)

type UpdateStatusOrderUseCase struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewUpdateStatusOrderUseCase(orderRepository repository.OrderRepositoryInterface) *UpdateStatusOrderUseCase {
	return &UpdateStatusOrderUseCase{orderRepository: orderRepository}
}

func (c UpdateStatusOrderUseCase) Execute(input dtos.InputUpdateOrderDto) error {
	order, err := c.orderRepository.FindById(input.ID)

	if err != nil {
		return err
	}

	err = order.UpdateStatus(input.Status)

	if err != nil {
		return err
	}

	err = c.orderRepository.UpdateStatus(*order)

	if err != nil {
		return err
	}

	return nil
}
