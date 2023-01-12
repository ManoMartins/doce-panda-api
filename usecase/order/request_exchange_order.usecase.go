package order

import (
	couponEntity "doce-panda/domain/coupon/entity"
	couponRepository "doce-panda/domain/coupon/repository"
	"doce-panda/domain/order/entity"
	"doce-panda/domain/order/repository"
	"doce-panda/domain/order/service"
	"doce-panda/usecase/order/dtos"
	"fmt"
)

type RequestExchangeOrderUseCase struct {
	orderRepository  repository.OrderRepositoryInterface
	couponRepository couponRepository.CouponRepositoryInterface
}

func NewRequestExchangeOrderUseCase(orderRepository repository.OrderRepositoryInterface, couponRepository couponRepository.CouponRepositoryInterface) *RequestExchangeOrderUseCase {
	return &RequestExchangeOrderUseCase{orderRepository: orderRepository, couponRepository: couponRepository}
}

func (c RequestExchangeOrderUseCase) Execute(input dtos.InputRequestExchangeOrderDto) (*dtos.OutputRequestExchangeOrderDto, error) {
	order, err := c.orderRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	var orderItemForExchange []entity.OrderItem
	for _, inputOrderItem := range input.OrderItems {
		hasProduct := false
		for _, orderItem := range order.OrderItems {
			if orderItem.ProductID == inputOrderItem.ProductID {
				if orderItem.Quantity < inputOrderItem.Quantity {
					return nil, fmt.Errorf("Quantidade não disponível para troca")
				}

				hasProduct = true
				orderItemForExchange = append(orderItemForExchange, entity.OrderItem{
					ID:           orderItem.ID,
					ProductID:    orderItem.ProductID,
					OrderID:      orderItem.OrderID,
					Quantity:     inputOrderItem.Quantity,
					TotalInCents: orderItem.Product.PriceInCents * inputOrderItem.Quantity,
					Product:      orderItem.Product,
					CreatedAt:    orderItem.CreatedAt,
					UpdatedAt:    orderItem.UpdatedAt,
				})
			}
		}

		if hasProduct == false {
			return nil, fmt.Errorf("Produto não permitido para troca")
		}
	}

	err = order.RequestExchange()

	if err != nil {
		return nil, err
	}

	if err = c.orderRepository.UpdateStatus(*order); err != nil {
		return nil, err
	}

	coupon, err := couponEntity.NewCoupon(couponEntity.Coupon{
		Amount: service.Total(orderItemForExchange),
		UserID: input.UserID,
	})

	if err != nil {
		return nil, err
	}

	if err = c.couponRepository.Create(*coupon); err != nil {
		return nil, err
	}

	return &dtos.OutputRequestExchangeOrderDto{
		CouponID:    coupon.ID,
		VoucherCode: coupon.VoucherCode,
		Amount:      coupon.Amount,
		Status:      coupon.Status,
	}, nil
}
