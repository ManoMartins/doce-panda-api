package handlers

import (
	"doce-panda/businessController/product"
	"doce-panda/businessController/product/dtos"
	"doce-panda/infra/db/gorm"
	"doce-panda/infra/gorm/product/repository"
	"github.com/gofiber/fiber/v2"
)

func FindAllProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		output, err := product.NewFindAllProductBusinessController(productRepository).Execute()

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

		input := dtos.InputFindProductDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		output, err := product.NewFindProductBusinessController(productRepository).Execute(input)

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
		body := new(dtos.InputCreateProductDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		input := dtos.InputCreateProductDto{
			Name:         body.Name,
			PriceInCents: body.PriceInCents,
			Description:  body.Description,
			Flavor:       body.Flavor,
			Quantity:     body.Quantity,
			CategoryID:   body.CategoryID,
		}

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)
		output, err := product.NewCreateProductBusinessController(productRepository).Execute(input)

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

func UpdateProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		body := new(dtos.InputUpdateProductDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)
		findProduct, err := product.NewFindProductBusinessController(productRepository).Execute(dtos.InputFindProductDto{ID: id})

		input := dtos.InputUpdateProductDto{
			ID:           id,
			Name:         body.Name,
			PriceInCents: body.PriceInCents,
			Description:  body.Description,
			Flavor:       body.Flavor,
			Quantity:     body.Quantity,
			Status:       dtos.StatusEnum(findProduct.Status),
		}

		output, err := product.NewUpdateProductBusinessController(productRepository).Execute(input)

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

		input := dtos.InputDeleteProductDto{ID: id}

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		err := product.NewDeleteProductBusinessController(productRepository).Execute(input)

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

		productEntity, err := product.NewUploadProductBusinessController(productRepository).Execute(dtos.InputUploadProductDto{
			ID:   id,
			File: file,
		})

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"data":    productEntity,
		})
	}
}

func DisableProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		err := product.NewDisableProductBusinessController(productRepository).Execute(dtos.InputDisableProductDto{
			ID: id,
		})

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}

		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"success": true,
		})
	}
}

func EnableProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		db := gorm.NewDb()
		defer db.Close()

		productRepository := repository.NewProductRepository(db)

		err := product.NewEnableProductBusinessController(productRepository).Execute(dtos.InputEnableProductDto{
			ID: id,
		})

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      true,
				"errorMessage": err.Error(),
				"error":        err,
			})
		}

		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"success": true,
		})
	}
}

func CreateCategory() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		body := new(dtos.InputCreateCategoryDto)
		err := ctx.BodyParser(body)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		input := dtos.InputCreateCategoryDto{Description: body.Description}

		db := gorm.NewDb()
		defer db.Close()

		categoryRepository := repository.NewCategoryRepository(db)

		category, err := product.NewCreateCategoryBusinessController(categoryRepository).Execute(input)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success":      false,
				"error":        err,
				"errorMessage": err.Error(),
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"data":    category,
		})
	}
}
