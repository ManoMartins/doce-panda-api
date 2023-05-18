package order

import (
	"doce-panda/businessController/order/dtos"
	"doce-panda/domain/order/repository"
)

type UpdateStatusOrderBusinessController struct {
	orderRepository repository.OrderRepositoryInterface
}

func NewUpdateStatusOrderBusinessController(orderRepository repository.OrderRepositoryInterface) *UpdateStatusOrderBusinessController {
	return &UpdateStatusOrderBusinessController{orderRepository: orderRepository}
}

func (c UpdateStatusOrderBusinessController) Execute(input dtos.InputUpdateOrderDto) error {
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
