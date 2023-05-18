package handlers

import (
	"doce-panda/businessController/order"
	"doce-panda/businessController/order/dtos"
	"doce-panda/infra/db/gorm"
	couponRepository "doce-panda/infra/gorm/coupon/repository"
	"doce-panda/infra/gorm/order/repository"
	productRepository "doce-panda/infra/gorm/product/repository"
	"github.com/gofiber/fiber/v2"
	"time"
)

func FindByIdOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := dtos.InputFindByIdOrderDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		orderRepo := repository.NewOrderRepository(db)

		output, err := order.NewFindByIdOrderBusinessController(orderRepo).Execute(input)

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

		output, err := order.NewFindAllOrderBusinessController(orderRepo).Execute()

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
			AddressID:    body.AddressID,
			UserID:       userId,
			OrderItems:   body.OrderItems,
			Payments:     body.Payments,
			VoucherCodes: body.VoucherCodes,
		}

		orderRepo := repository.NewOrderRepository(db)
		couponRepo := couponRepository.NewCouponRepository(db)
		productRepo := productRepository.NewProductRepository(db)

		output, err := order.NewCreateOrderBusinessController(orderRepo, couponRepo, productRepo).Execute(input)

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

		err = order.NewUpdateStatusOrderBusinessController(orderRepo).Execute(input)

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
		orderPendingCouponRepo := repository.NewOrderPendingCouponRepository(db)

		output, err := order.NewRequestExchangeOrderBusinessController(orderRepo, couponRepo, orderPendingCouponRepo).Execute(input)

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

		input := dtos.InputAcceptExchangeRequestOrderDto{
			ID: id,
		}

		db := gorm.NewDb()
		defer db.Close()

		orderRepo := repository.NewOrderRepository(db)

		err := order.NewAcceptExchangeRequestOrderBusinessController(orderRepo).Execute(input)

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

func ExchangeReceivedOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := dtos.InputExchangeReceivedOrderDto{
			ID: id,
		}

		db := gorm.NewDb()
		defer db.Close()

		orderRepo := repository.NewOrderRepository(db)
		couponRepo := couponRepository.NewCouponRepository(db)
		orderPendingCouponRepo := repository.NewOrderPendingCouponRepository(db)

		err := order.NewExchangeReceivedOrderBusinessController(orderRepo, couponRepo, orderPendingCouponRepo).Execute(input)

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

		input := dtos.InputDenyRequestExchangeOrderDto{
			ID: id,
		}

		db := gorm.NewDb()
		defer db.Close()

		orderRepo := repository.NewOrderRepository(db)
		couponRepo := couponRepository.NewCouponRepository(db)
		orderPendingCouponRepo := repository.NewOrderPendingCouponRepository(db)

		err := order.NewDenyRequestExchangeOrderBusinessController(orderRepo, couponRepo, orderPendingCouponRepo).Execute(input)

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

func FindAllOrderByUserId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("userId").(string)

		input := dtos.InputFindAllOrderByUserIdDto{
			UserID: userId,
		}

		db := gorm.NewDb()
		defer db.Close()

		orderRepo := repository.NewOrderRepository(db)

		output, err := order.NewFindAllOrderByUserIdBusinessController(orderRepo).Execute(input)

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

func ReportOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		query := new(struct {
			StartDate string
			EndDate   string
		})
		err := ctx.QueryParser(query)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		startDate, _ := time.Parse("2006-01-02", query.StartDate)
		endDate, _ := time.Parse("2006-01-02", query.EndDate)

		db := gorm.NewDb()
		defer db.Close()

		input := dtos.InputReportOrderDto{
			StartDate: startDate,
			EndDate:   endDate,
		}

		orderRepo := repository.NewOrderRepository(db)

		output, err := order.NewReportOrderBusinessController(orderRepo).Execute(input)

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
