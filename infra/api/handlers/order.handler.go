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

		orderRepo := repository.NewOrderRepository(db)

		output, err := order.NewFindByIdOrderUseCase(orderRepo).Execute(input)

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

		orderRepo := repository.NewOrderRepository(db)

		output, err := order.NewFindAllOrderUseCase(orderRepo).Execute()

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
		userId := ctx.Locals("userId").(string)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		db := gorm.NewDb()
		defer db.Close()

		input := dtos.InputCreateOrderDto{
			AddressID:  body.AddressID,
			UserID:     userId,
			OrderItems: body.OrderItems,
			Payments:   body.Payments,
		}

		orderRepo := repository.NewOrderRepository(db)
		productRepo := productRepository.NewProductRepository(db)

		output, err := order.NewCreateOrderUseCase(orderRepo, productRepo).Execute(input)

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
