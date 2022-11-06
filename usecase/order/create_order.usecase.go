package order

import (
	"doce-panda/domain/order/entity"
	"doce-panda/domain/order/repository"
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

	totalInCents := 0
	for _, orderItem := range orderItems {
		totalInCents += orderItem.TotalInCents
	}

	order, err := entity.NewOrder(entity.Order{
		TotalInCents: totalInCents,
	})

	order.AddOrderItems(orderItems)

	if err != nil {
		return nil, err
	}

	c.orderRepository.Create(*order)

	for _, orderItem := range order.OrderItems {
		err := orderItem.Product.RemoveQuantity(orderItem.Quantity)

		if err != nil {
			return nil, err
		}

		c.productRepository.Update(orderItem.Product)
	}

	return &dtos.OutputCreateOrderDto{
		ID:           order.ID,
		OrderItems:   order.OrderItems,
		TotalInCents: order.TotalInCents,
		Status:       order.Status,
	}, nil
}
