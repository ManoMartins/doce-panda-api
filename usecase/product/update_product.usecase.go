package product

import (
	"doce-panda/domain/product/entity"
	"doce-panda/domain/product/repository"
	"doce-panda/usecase/product/dtos"
)

type UpdateProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewUpdateProductUseCase(productRepository repository.ProductRepositoryInterface) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepository: productRepository,
	}
}

func (c UpdateProductUseCase) Execute(input dtos.InputUpdateProductDto) (*dtos.OutputUpdateProductDto, error) {
	product, err := c.productRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	err = product.Update(entity.ProductUpdateProps{
		Name:         input.Name,
		PriceInCents: input.PriceInCents,
		Status:       entity.StatusEnum(input.Status),
		Description:  input.Description,
		Flavor:       input.Flavor,
		Quantity:     input.Quantity,
	})

	if err != nil {
		return nil, err
	}

	err = c.productRepository.Update(*product)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputUpdateProductDto{
		ID:           product.ID,
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Description:  product.Description,
		Flavor:       product.Flavor,
		Quantity:     product.Quantity,
		ImageUrl:     product.ImageUrl,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}, nil
}
