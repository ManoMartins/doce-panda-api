package order

import (
	couponRepository "doce-panda/domain/coupon/repository"
	"doce-panda/domain/order/repository"
	"doce-panda/usecase/order/dtos"
)

type AcceptRequestExchangeOrderUseCase struct {
	orderRepository  repository.OrderRepositoryInterface
	couponRepository couponRepository.CouponRepositoryInterface
}

func NewAcceptRequestExchangeOrderUseCase(orderRepository repository.OrderRepositoryInterface, couponRepository couponRepository.CouponRepositoryInterface) *AcceptRequestExchangeOrderUseCase {
	return &AcceptRequestExchangeOrderUseCase{orderRepository: orderRepository, couponRepository: couponRepository}
}

func (c AcceptRequestExchangeOrderUseCase) Execute(input dtos.InputAcceptRequestExchangeOrderDto) error {
	order, err := c.orderRepository.FindById(input.ID)

	if err != nil {
		return err
	}

	coupon, err := c.couponRepository.FindByVoucherCode(input.VoucherCode)

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

	err = coupon.AcceptToUseCoupon()

	if err != nil {
		return err
	}

	err = c.couponRepository.UpdateStatus(*coupon)

	if err != nil {
		return err
	}

	return nil
}
