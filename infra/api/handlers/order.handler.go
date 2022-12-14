package handlers

import (
	"doce-panda/infra/db/gorm"
	couponRepository "doce-panda/infra/gorm/coupon/repository"
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
			AddressID:   body.AddressID,
			UserID:      userId,
			OrderItems:  body.OrderItems,
			Payments:    body.Payments,
			VoucherCode: body.VoucherCode,
		}

		orderRepo := repository.NewOrderRepository(db)
		couponRepo := couponRepository.NewCouponRepository(db)
		productRepo := productRepository.NewProductRepository(db)

		output, err := order.NewCreateOrderUseCase(orderRepo, couponRepo, productRepo).Execute(input)

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

func UpdateOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		body := new(dtos.InputUpdateOrderDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		input := dtos.InputUpdateOrderDto{ID: id, Status: body.Status}

		db := gorm.NewDb()
		defer db.Close()

		orderRepo := repository.NewOrderRepository(db)

		err = order.NewUpdateStatusOrderUseCase(orderRepo).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
		})
	}
}

func RequestExchangeOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		userId := ctx.Locals("userId").(string)
		body := new(dtos.InputRequestExchangeOrderDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		input := dtos.InputRequestExchangeOrderDto{
			ID:         id,
			UserID:     userId,
			OrderItems: body.OrderItems,
		}

		db := gorm.NewDb()
		defer db.Close()

		orderRepo := repository.NewOrderRepository(db)
		couponRepo := couponRepository.NewCouponRepository(db)

		output, err := order.NewRequestExchangeOrderUseCase(orderRepo, couponRepo).Execute(input)

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

func AcceptRequestExchangeOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		body := new(dtos.InputAcceptRequestExchangeOrderDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		input := dtos.InputAcceptRequestExchangeOrderDto{
			ID:          id,
			VoucherCode: body.VoucherCode,
		}

		db := gorm.NewDb()
		defer db.Close()

		orderRepo := repository.NewOrderRepository(db)
		couponRepo := couponRepository.NewCouponRepository(db)

		err = order.NewAcceptRequestExchangeOrderUseCase(orderRepo, couponRepo).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
		})
	}
}

func DenyRequestExchangeOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		body := new(dtos.InputDenyRequestExchangeOrderDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		input := dtos.InputDenyRequestExchangeOrderDto{
			ID:          id,
			VoucherCode: body.VoucherCode,
		}

		db := gorm.NewDb()
		defer db.Close()

		orderRepo := repository.NewOrderRepository(db)
		couponRepo := couponRepository.NewCouponRepository(db)

		err = order.NewDenyRequestExchangeOrderUseCase(orderRepo, couponRepo).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
		})
	}
}
