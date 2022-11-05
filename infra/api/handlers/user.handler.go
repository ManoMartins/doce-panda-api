package handlers

import (
	"doce-panda/infra/db/gorm"
	"doce-panda/infra/gorm/user/repository"
	"doce-panda/usecase/user"
	"doce-panda/usecase/user/dtos"
	"github.com/gofiber/fiber/v2"
)

func FindUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := dtos.InputFindUserDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		userRepository := repository.NewUserRepository(db)

		output, err := user.NewFindUserUseCase(userRepository).Execute(input)

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

func FindAllUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		db := gorm.NewDb()
		defer db.Close()

		userRepository := repository.NewUserRepository(db)

		output, err := user.NewFindAllUserUseCase(userRepository).Execute()

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    output,
		})
	}
}

func CreateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		body := new(dtos.InputCreateUserDto)
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

		userRepository := repository.NewUserRepository(db)
		output, err := user.NewCreateUserUseCase(userRepository).Execute(*body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"data":    output,
		})
	}
}

func UpdateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		body := new(dtos.InputUpdateUserDto)
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
		userRepository := repository.NewUserRepository(db)

		input := dtos.InputUpdateUserDto{
			ID:             id,
			Name:           body.Name,
			Email:          body.Email,
			Gender:         body.Gender,
			PhoneNumber:    body.PhoneNumber,
			DocumentNumber: body.DocumentNumber,
		}

		output, err := user.NewUpdateUserUseCase(userRepository).Execute(input)

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

func DeleteUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := dtos.InputDestroyUserDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		userRepository := repository.NewUserRepository(db)

		err := user.NewDestroyUserUseCase(userRepository).Execute(input)

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

func CreateAddress() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		body := new(dtos.InputCreateAddressDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		body.UserID = id

		db := gorm.NewDb()
		defer db.Close()

		addressRepository := repository.NewAddressRepository(db)
		output, err := user.NewCreateAddressUseCase(addressRepository).Execute(*body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"data":    output,
		})
	}
}

func DeleteAddress() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		addressId := ctx.Params("addressId")

		input := dtos.InputDestroyAddressDto{
			ID:        id,
			AddressID: addressId,
		}

		db := gorm.NewDb()
		defer db.Close()

		addressRepository := repository.NewAddressRepository(db)

		err := user.NewDestroyAddressUseCase(addressRepository).Execute(input)

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
