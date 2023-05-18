package handlers

import (
	"doce-panda/businessController/user"
	"doce-panda/businessController/user/dtos"
	"doce-panda/infra/db/gorm"
	"doce-panda/infra/gorm/user/repository"
	"github.com/gofiber/fiber/v2"
)

func LoginUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		body := new(dtos.InputAuthenticationUserDto)
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

		userRepo := repository.NewUserRepository(db)
		output, err := user.NewAuthenticateUserBusinessController(userRepo).Execute(*body)

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

func FindUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := dtos.InputFindUserDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		userRepository := repository.NewUserRepository(db)

		output, err := user.NewFindUserBusinessController(userRepository).Execute(input)

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

		output, err := user.NewFindAllUserBusinessController(userRepository).Execute()

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
		output, err := user.NewCreateUserBusinessController(userRepository).Execute(*body)

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
			Gender:         body.Gender,
			PhoneNumber:    body.PhoneNumber,
			DocumentNumber: body.DocumentNumber,
		}

		output, err := user.NewUpdateUserBusinessController(userRepository).Execute(input)

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

		err := user.NewDestroyUserBusinessController(userRepository).Execute(input)

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
		id := ctx.Locals("userId").(string)
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
		output, err := user.NewCreateAddressBusinessController(addressRepository).Execute(*body)

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

func FindAllAddressByUserId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Locals("userId").(string)

		db := gorm.NewDb()
		defer db.Close()

		input := dtos.InputFindAllAddressDto{UserID: id}

		addressRepository := repository.NewAddressRepository(db)
		output, err := user.NewFindAllAddressBusinessController(addressRepository).Execute(input)

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

func UpdateAddress() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Locals("userId").(string)
		addressId := ctx.Params("addressId")

		body := new(dtos.InputUpdateAddressDto)
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
		addressRepository := repository.NewAddressRepository(db)

		input := dtos.InputUpdateAddressDto{
			ID:           addressId,
			City:         body.City,
			State:        body.State,
			Street:       body.Street,
			Number:       body.Number,
			ZipCode:      body.ZipCode,
			Neighborhood: body.Neighborhood,
			IsMain:       body.IsMain,
			UserID:       id,
		}

		output, err := user.NewUpdateAddressBusinessController(addressRepository).Execute(input)

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

func DeleteAddress() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Locals("userId").(string)
		addressId := ctx.Params("addressId")

		input := dtos.InputDestroyAddressDto{
			ID:        id,
			AddressID: addressId,
		}

		db := gorm.NewDb()
		defer db.Close()

		addressRepository := repository.NewAddressRepository(db)

		err := user.NewDestroyAddressBusinessController(addressRepository).Execute(input)

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

func FindByIdAddress() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		addressId := ctx.Params("addressId")

		input := dtos.InputFindByIdAddressDto{
			ID: addressId,
		}

		db := gorm.NewDb()
		defer db.Close()

		addressRepository := repository.NewAddressRepository(db)

		output, err := user.NewFindByIdAddressBusinessController(addressRepository).Execute(input)

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
