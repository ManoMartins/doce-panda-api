package order

import (
	"doce-panda/domain/order/repository"
	"doce-panda/usecase/order/dtos"
)

type AcceptExchangeRequestOrderUseCase struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewAcceptExchangeRequestOrderUseCase(orderRepository repository.OrderRepositoryInterface) *AcceptExchangeRequestOrderUseCase {
	return &AcceptExchangeRequestOrderUseCase{orderRepository: orderRepository}
}

func (c AcceptExchangeRequestOrderUseCase) Execute(input dtos.InputAcceptExchangeRequestOrderDto) error {
	order, err := c.orderRepository.FindById(input.ID)

	if err != nil {
		return err
	}

	err = order.AcceptExchangeRequest()

	if err != nil {
		return err
	}

	err = c.orderRepository.UpdateStatus(*order)

	if err != nil {
		return err
	}

	return nil
}
