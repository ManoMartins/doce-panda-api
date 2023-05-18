package handlers

import (
	"doce-panda/businessController/payment"
	"doce-panda/businessController/payment/dtos"
	"doce-panda/infra/db/gorm"
	"doce-panda/infra/gorm/payment/repository"
	"github.com/gofiber/fiber/v2"
)

func FindByIdCreditCard() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := dtos.InputFindByIdCreditCardDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		creditCardRepository := repository.NewCreditCardRepository(db)

		output, err := payment.NewFindByIdCreditCardBusinessController(creditCardRepository).Execute(input)

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

func FindAllCreditCard() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		db := gorm.NewDb()
		defer db.Close()

		creditCardRepository := repository.NewCreditCardRepository(db)

		output, err := payment.NewFindAllCreditCardBusinessController(creditCardRepository).Execute()

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

func CreateCreditCard() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		body := new(dtos.InputCreateCreditCardDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		db := gorm.NewDb()
		defer db.Close()

		input := dtos.InputCreateCreditCardDto{
			CardLastNumber:     body.CardLastNumber,
			CardHolder:         body.CardHolder,
			CardIdentification: body.CardIdentification,
			CardSecurityCode:   body.CardSecurityCode,
			CardExpirationDate: body.CardExpirationDate,
			CardBrand:          body.CardBrand,
		}

		creditCardRepository := repository.NewCreditCardRepository(db)

		output, err := payment.NewCreateCreditCardBusinessController(creditCardRepository).Execute(input)

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

func DeleteCreditCard() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := dtos.InputDeleteCreditCardDto{
			ID: id,
		}

		db := gorm.NewDb()
		defer db.Close()

		creditCardRepository := repository.NewCreditCardRepository(db)

		err := payment.NewDeleteCreditCardBusinessController(creditCardRepository).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.SendStatus(fiber.StatusNoContent)
	}
}
