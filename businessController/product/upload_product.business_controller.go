package product

import (
	"doce-panda/businessController/product/dtos"
	"doce-panda/domain/product/repository"
	"fmt"
	"io"
	"os"
)

type UploadProductBusinessController struct {
	productRepository repository.ProductRepositoryInterface
}

func NewUploadProductBusinessController(productRepository repository.ProductRepositoryInterface) *UploadProductBusinessController {
	return &UploadProductBusinessController{
		productRepository: productRepository,
	}
}

func (c UploadProductBusinessController) Execute(input dtos.InputUploadProductDto) (*dtos.OutputUploadProductDto, error) {
	product, err := c.productRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	imageUri := fmt.Sprintf("/uploads/%s", input.File.Filename)

	src, err := input.File.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(fmt.Sprintf("./tmp/%s", input.File.Filename))
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	product.AddImageUrl(imageUri)

	err = c.productRepository.Update(*product)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputUploadProductDto{
		ID:           product.ID,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Description:  product.Description,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		ImageUrl:     "http://localhost:3333" + product.ImageUrl,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}, nil
}
