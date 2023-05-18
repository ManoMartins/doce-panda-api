package order

import (
	"doce-panda/businessController/order/dtos"
	"doce-panda/domain/order/repository"
)

type AcceptExchangeRequestOrderBusinessController struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewAcceptExchangeRequestOrderBusinessController(
	orderRepository repository.OrderRepositoryInterface,
) *AcceptExchangeRequestOrderBusinessController {
	return &AcceptExchangeRequestOrderBusinessController{
		orderRepository: orderRepository,
	}
}

func (c AcceptExchangeRequestOrderBusinessController) Execute(input dtos.InputAcceptExchangeRequestOrderDto) error {
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
