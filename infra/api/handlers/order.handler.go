package handlers

import (
	"doce-panda/infra/db/gorm"
	"doce-panda/infra/gorm/order/repository"
	productRepository "doce-panda/infra/gorm/product/repository"
	"doce-panda/usecase/order"
	"doce-panda/usecase/order/dtos"
	"github.com/gofiber/fiber/v2"
)

func FindByIdOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := dtos.InputFindByIdOrderDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		orderRepository := repository.NewOrderRepository(db)

		output, err := order.NewFindByIdOrderUseCase(orderRepository).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"data":    output,
		})

	}
}
func FindAllOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		db := gorm.NewDb()
		defer db.Close()

		orderRepository := repository.NewOrderRepository(db)

		output, err := order.NewFindAllOrderUseCase(orderRepository).Execute()

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"data":    output,
		})

	}
}

func CreateOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		body := new(dtos.InputCreateOrderDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		db := gorm.NewDb()
		defer db.Close()

		input := dtos.InputCreateOrderDto{
			OrderItems: body.OrderItems,
		}

		orderRepository := repository.NewOrderRepository(db)
		productRepository := productRepository.NewProductRepository(db)

		output, err := order.NewCreateOrderUseCase(orderRepository, productRepository).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"data":    output,
		})
	}
}
