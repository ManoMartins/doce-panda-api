package order

import (
	"doce-panda/businessController/order/dtos"
	couponRepository "doce-panda/domain/coupon/repository"
	"doce-panda/domain/order/repository"
)

type DenyRequestExchangeOrderBusinessController struct {
	orderRepository              repository.OrderRepositoryInterface
	couponRepository             couponRepository.CouponRepositoryInterface
	orderPendingCouponRepository repository.OrderPendingCouponRepositoryInterface
}

func NewDenyRequestExchangeOrderBusinessController(
	orderRepository repository.OrderRepositoryInterface,
	couponRepository couponRepository.CouponRepositoryInterface,
	orderPendingCouponRepository repository.OrderPendingCouponRepositoryInterface,
) *DenyRequestExchangeOrderBusinessController {
	return &DenyRequestExchangeOrderBusinessController{
		orderRepository:              orderRepository,
		couponRepository:             couponRepository,
		orderPendingCouponRepository: orderPendingCouponRepository,
	}
}

func (c DenyRequestExchangeOrderBusinessController) Execute(input dtos.InputDenyRequestExchangeOrderDto) error {
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

	err = c.orderPendingCouponRepository.DeleteByOrderID(input.ID)

	if err != nil {
		return err
	}

	return nil
}
