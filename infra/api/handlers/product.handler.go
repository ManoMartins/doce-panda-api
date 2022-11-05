package handlers

import (
	"doce-panda/infra/db/gorm"
	"doce-panda/infra/gorm/product/repository"
	"doce-panda/usecase/product/create"
	"doce-panda/usecase/product/delete"
	"doce-panda/usecase/product/disable"
	"doce-panda/usecase/product/enable"
	"doce-panda/usecase/product/find"
	"doce-panda/usecase/product/find_all"
	"doce-panda/usecase/product/update"
	"doce-panda/usecase/product/upload"
	"github.com/gofiber/fiber/v2"
)

func FindAllProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		output, err := find_all.NewFindAllProductUseCase(productRepository).Execute()

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    output,
		})
	}
}
func FindProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := find.InputFindProductDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		output, err := find.NewFindProductUseCase(productRepository).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    output,
		})
	}
}

func CreateProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		body := new(create.InputCreateProductDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		input := create.InputCreateProductDto{
			Name:         body.Name,
			PriceInCents: body.PriceInCents,
			Description:  body.Description,
			Flavor:       body.Flavor,
			Quantity:     body.Quantity,
		}

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)
		output, err := create.NewCreateProductUseCase(productRepository).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"data": fiber.Map{
					"message": err.Error(),
				},
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"data":    output,
		})
	}
}

func UpdateProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		body := new(update.InputUpdateProductDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)
		findProduct, err := find.NewFindProductUseCase(productRepository).Execute(find.InputFindProductDto{ID: id})

		input := update.InputUpdateProductDto{
			ID:           id,
			Name:         body.Name,
			PriceInCents: body.PriceInCents,
			Description:  body.Description,
			Flavor:       body.Flavor,
			Quantity:     body.Quantity,
			Status:       update.StatusEnum(findProduct.Status),
		}

		output, err := update.NewUpdateProductUseCase(productRepository).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"data": fiber.Map{
					"message": err.Error(),
				},
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"data":    output,
		})
	}
}

func DestroyProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		input := destroy.InputDeleteProductDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		err := destroy.NewDeleteProductUseCase(productRepository).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"success": true,
		})
	}
}

func UploadProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		file, err := ctx.FormFile("file")

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		product, err := upload.NewUploadProductUseCase(productRepository).Execute(upload.InputUploadProductDto{
			ID:   id,
			File: file,
		})

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"data":    product,
		})
	}
}

func DisableProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		product, err := disable.NewDisableProductUseCase(productRepository).Execute(disable.InputDisableProductDto{
			ID: id,
		})

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"data":    product,
		})
	}
}

func EnableProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		product, err := enable.NewEnableProductUseCase(productRepository).Execute(enable.InputEnableProductDto{
			ID: id,
		})

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      true,
				"errorMessage": err.Error(),
				"error":        err,
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"data":    product,
		})
	}
}
