package upload

import (
	"doce-panda/domain/product/repository"
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

func (c UploadProductUseCase) Execute(input InputUploadProductDto) (*OutputUploadProductDto, error) {

	imageUri := fmt.Sprintf("/uploads/products/%s", input.File.Filename)

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

	productUpdated, err := c.productRepository.Upload(input.ID, imageUri)

	if err != nil {
		return nil, err
	}

	return &OutputUploadProductDto{
		ID:           productUpdated.ID,
		Name:         productUpdated.Name,
		PriceInCents: productUpdated.PriceInCents,
		Description:  productUpdated.Description,
		Flavor:       productUpdated.Flavor,
		Quantity:     productUpdated.Quantity,
		ImageUrl:     productUpdated.ImageUrl,
		CreatedAt:    productUpdated.CreatedAt,
		UpdatedAt:    productUpdated.UpdatedAt,
	}, nil
}
