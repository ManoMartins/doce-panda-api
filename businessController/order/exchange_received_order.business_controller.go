package order

import (
	"doce-panda/businessController/order/dtos"
	couponRepository "doce-panda/domain/coupon/repository"
	"doce-panda/domain/order/repository"
)

type ExchangeReceivedOrderBusinessController struct {
	orderRepository              repository.OrderRepositoryInterface
	couponRepository             couponRepository.CouponRepositoryInterface
	orderPendingCouponRepository repository.OrderPendingCouponRepositoryInterface
}

func NewExchangeReceivedOrderBusinessController(
	orderRepository repository.OrderRepositoryInterface,
	couponRepository couponRepository.CouponRepositoryInterface,
	orderPendingCouponRepository repository.OrderPendingCouponRepositoryInterface,
) *ExchangeReceivedOrderBusinessController {
	return &ExchangeReceivedOrderBusinessController{
		orderRepository:              orderRepository,
		couponRepository:             couponRepository,
		orderPendingCouponRepository: orderPendingCouponRepository,
	}
}

func (c ExchangeReceivedOrderBusinessController) Execute(input dtos.InputExchangeReceivedOrderDto) error {
	order, err := c.orderRepository.FindById(input.ID)

	if err != nil {
		return err
	}

	orderPendingCoupon, err := c.orderPendingCouponRepository.FindByOrderId(input.ID)

	if err != nil {
		return err
	}

	coupon, err := c.couponRepository.FindById(orderPendingCoupon.CouponID)

	if err != nil {
		return err
	}

	err = order.ExchangeReceived()

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

	err = c.orderPendingCouponRepository.DeleteByOrderID(input.ID)

	if err != nil {
		return err
	}

	return nil
}
