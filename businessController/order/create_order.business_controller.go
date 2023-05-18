package order

import (
	"doce-panda/businessController/order/dtos"
	couponEntity "doce-panda/domain/coupon/entity"
	couponRepository "doce-panda/domain/coupon/repository"
	"doce-panda/domain/order/entity"
	"doce-panda/domain/order/repository"
	"doce-panda/domain/order/service"
	paymentEntity "doce-panda/domain/payment/entity"
	paymentService "doce-panda/domain/payment/service"
	productEntity "doce-panda/domain/product/entity"
	productRepository "doce-panda/domain/product/repository"
	"fmt"
)

type CreateOrderBusinessController struct {
	orderRepository   repository.OrderRepositoryInterface
	couponRepository  couponRepository.CouponRepositoryInterface
	productRepository productRepository.ProductRepositoryInterface
}

func NewCreateOrderBusinessController(orderRepository repository.OrderRepositoryInterface, couponRepository couponRepository.CouponRepositoryInterface, productRepository productRepository.ProductRepositoryInterface) *CreateOrderBusinessController {
	return &CreateOrderBusinessController{orderRepository: orderRepository, couponRepository: couponRepository, productRepository: productRepository}
}

func (c CreateOrderBusinessController) Execute(input dtos.InputCreateOrderDto) (*dtos.OutputCreateOrderDto, error) {
	var orderItems []entity.OrderItem
	for _, inputOrderItem := range input.OrderItems {
		product, err := c.productRepository.FindById(inputOrderItem.ProductID)

		if err != nil {
			return nil, err
		}

		if product.Status == productEntity.DISABLED {
			return nil, fmt.Errorf("Não foi possível comprar, pois o produto está desabilitado.")
		}

		if product.Quantity <= 0 || product.Quantity < inputOrderItem.Quantity {
			return nil, fmt.Errorf("Não foi possível comprar, pois não tem quantidade em estoque.")
		}

		orderItem, err := entity.NewOrderItem(entity.OrderItem{
			ProductID:    inputOrderItem.ProductID,
			Quantity:     inputOrderItem.Quantity,
			TotalInCents: product.PriceInCents * inputOrderItem.Quantity,
			Product:      *product,
		})

		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, *orderItem)
	}

	var paymentsEntity []paymentEntity.CreditCard
	for _, payment := range input.Payments {
		paymentsEntity = append(paymentsEntity, paymentEntity.CreditCard{ID: payment.PaymentID, TotalInCents: payment.TotalInCents})
	}

	var coupons []couponEntity.Coupon
	for _, voucherCode := range input.VoucherCodes {
		coupon, err := c.couponRepository.FindByVoucherCode(voucherCode)

		if err != nil {
			return nil, err
		}

		coupons = append(coupons, *coupon)
	}

	paymentTotalInCents := paymentService.Total(paymentsEntity)

	order, err := entity.NewOrder(entity.Order{
		SubTotalInCents: service.Total(orderItems),
		Payments:        paymentsEntity,
		DeliveredFee:    500,
		UserID:          input.UserID,
		Coupons:         coupons,
		AddressID:       input.AddressID,
	})

	if err != nil {
		return nil, err
	}

	var moneyExchange int
	if len(coupons) != 0 {
		moneyExchange, err = order.ApplyCoupons(coupons)
	}

	if err != nil {
		return nil, err
	}

	if paymentTotalInCents != (order.TotalInCents) {
		return nil, fmt.Errorf("O valor do pagamento está diferente do total")
	}

	order.AddOrderItems(orderItems)

	if err = c.orderRepository.Create(*order); err != nil {
		return nil, err
	}

	for _, orderItem := range order.OrderItems {
		if err := orderItem.Product.RemoveQuantity(orderItem.Quantity); err != nil {
			return nil, err
		}

		if err = c.productRepository.Update(orderItem.Product); err != nil {
			return nil, err
		}
	}

	var couponMoneyExchange *couponEntity.Coupon
	if len(coupons) != 0 {
		for _, coupon := range coupons {

			if err = coupon.UseCoupon(); err != nil {
				return nil, err
			}

			if err = c.couponRepository.UpdateStatus(coupon); err != nil {
				return nil, err
			}
		}

		if moneyExchange > 0 {
			couponMoneyExchange, err = couponEntity.NewCoupon(couponEntity.Coupon{
				UserID: input.UserID,
				Amount: moneyExchange,
				Status: couponEntity.NEW,
			})

			if err != nil {
				return nil, err
			}

			c.couponRepository.Create(*couponMoneyExchange)
		}
	}

	return &dtos.OutputCreateOrderDto{
		ID:                  order.ID,
		OrderItems:          order.OrderItems,
		TotalInCents:        order.TotalInCents,
		Status:              order.Status,
		CouponMoneyExchange: couponMoneyExchange,
	}, nil
}
