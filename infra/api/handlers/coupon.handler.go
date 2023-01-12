package handlers

import (
	"doce-panda/infra/db/gorm"
	"doce-panda/infra/gorm/coupon/repository"
	"doce-panda/usecase/coupon"
	"doce-panda/usecase/coupon/dtos"
	"github.com/gofiber/fiber/v2"
)

func FindAllCoupon() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		db := gorm.NewDb()
		defer db.Close()

		couponRepo := repository.NewCouponRepository(db)

		output, err := coupon.NewFindAllCouponUseCase(couponRepo).Execute()

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

func FindByIdCoupon() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := dtos.InputFindByIdCouponDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		couponRepo := repository.NewCouponRepository(db)

		output, err := coupon.NewFindByIdCouponUseCase(couponRepo).Execute(input)

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

func UpdateStatusCoupon() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		userId := ctx.Locals("userId").(string)
		body := new(dtos.InputUpdateStatusCouponDto)

		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		input := dtos.InputUpdateStatusCouponDto{
			ID:     id,
			UserID: userId,
			Status: body.Status,
		}

		db := gorm.NewDb()
		defer db.Close()

		couponRepo := repository.NewCouponRepository(db)

		err = coupon.NewUpdateStatusCouponUseCase(couponRepo).Execute(input)

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
