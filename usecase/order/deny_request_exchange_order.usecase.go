package order

import (
	couponRepository "doce-panda/domain/coupon/repository"
	"doce-panda/domain/order/repository"
	"doce-panda/usecase/order/dtos"
)

type DenyRequestExchangeOrderUseCase struct {
	orderRepository  repository.OrderRepositoryInterface
	couponRepository couponRepository.CouponRepositoryInterface
}

func NewDenyRequestExchangeOrderUseCase(orderRepository repository.OrderRepositoryInterface, couponRepository couponRepository.CouponRepositoryInterface) *DenyRequestExchangeOrderUseCase {
	return &DenyRequestExchangeOrderUseCase{orderRepository: orderRepository, couponRepository: couponRepository}
}

func (c DenyRequestExchangeOrderUseCase) Execute(input dtos.InputDenyRequestExchangeOrderDto) error {
	order, err := c.orderRepository.FindById(input.ID)

	if err != nil {
		return err
	}

	coupon, err := c.couponRepository.FindByVoucherCode(input.VoucherCode)

	if err != nil {
		return err
	}

	err = order.DenyExchangeRequest()

	if err != nil {
		return err
	}

	err = c.orderRepository.UpdateStatus(*order)

	if err != nil {
		return err
	}

	err = c.couponRepository.Delete(coupon.ID)

	if err != nil {
		return err
	}

	return nil
}
