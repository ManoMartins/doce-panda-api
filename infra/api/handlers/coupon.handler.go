package handlers

import (
	"doce-panda/businessController/coupon"
	"doce-panda/businessController/coupon/dtos"
	"doce-panda/infra/db/gorm"
	"doce-panda/infra/gorm/coupon/repository"
	"github.com/gofiber/fiber/v2"
)

func FindAllCoupon() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		db := gorm.NewDb()
		defer db.Close()

		couponRepo := repository.NewCouponRepository(db)

		output, err := coupon.NewFindAllCouponBusinessController(couponRepo).Execute()

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

		output, err := coupon.NewFindByIdCouponBusinessController(couponRepo).Execute(input)

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

		err = coupon.NewUpdateStatusCouponBusinessController(couponRepo).Execute(input)

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

func FindByVoucherCodeCoupon() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		voucherCode := ctx.Params("voucherCode")

		input := dtos.InputFindByVoucherCodeCouponDto{VoucherCode: voucherCode}

		db := gorm.NewDb()
		defer db.Close()

		couponRepo := repository.NewCouponRepository(db)
		output, err := coupon.NewFindByVoucherCodeCouponBusinessController(couponRepo).Execute(input)

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
