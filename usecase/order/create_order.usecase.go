package order

import (
	"doce-panda/domain/order/entity"
	"doce-panda/domain/order/repository"
	"doce-panda/domain/order/service"
	paymentEntity "doce-panda/domain/payment/entity"
	paymentService "doce-panda/domain/payment/service"
	productEntity "doce-panda/domain/product/entity"
	productRepository "doce-panda/domain/product/repository"
	"doce-panda/usecase/order/dtos"
	"fmt"
)

type CreateOrderUseCase struct {
	orderRepository   repository.OrderRepositoryInterface
	productRepository productRepository.ProductRepositoryInterface
}

func NewCreateOrderUseCase(orderRepository repository.OrderRepositoryInterface, productRepository productRepository.ProductRepositoryInterface) *CreateOrderUseCase {
	return &CreateOrderUseCase{orderRepository: orderRepository, productRepository: productRepository}
}

func (c CreateOrderUseCase) Execute(input dtos.InputCreateOrderDto) (*dtos.OutputCreateOrderDto, error) {
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

	paymentTotalInCents := paymentService.Total(paymentsEntity)

	order, err := entity.NewOrder(entity.Order{
		TotalInCents: service.Total(orderItems),
		Payments:     paymentsEntity,
		UserID:       input.UserID,
		AddressID:    input.AddressID,
	})

	if err != nil {
		return nil, err
	}

	if paymentTotalInCents != (order.TotalInCents + 500) {
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

	return &dtos.OutputCreateOrderDto{
		ID:           order.ID,
		OrderItems:   order.OrderItems,
		TotalInCents: order.TotalInCents,
		Status:       order.Status,
	}, nil
}
