package product

import (
	"doce-panda/domain/product/repository"
	"doce-panda/usecase/product/dtos"
	"fmt"
	"io"
	"os"
)

type UploadProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewUploadProductUseCase(productRepository repository.ProductRepositoryInterface) *UploadProductUseCase {
	return &UploadProductUseCase{
		productRepository: productRepository,
	}
}

func (c UploadProductUseCase) Execute(input dtos.InputUploadProductDto) (*dtos.OutputUploadProductDto, error) {
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
