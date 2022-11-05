package update

import (
	"doce-panda/domain/product/entity"
	"doce-panda/domain/product/repository"
)

type UpdateProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewUpdateProductUseCase(productRepository repository.ProductRepositoryInterface) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepository: productRepository,
	}
}

func (c UpdateProductUseCase) Execute(input InputUpdateProductDto) (*OutputUpdateProductDto, error) {
	product := entity.Product{
		ID:           input.ID,
		Name:         input.Name,
		PriceInCents: input.PriceInCents,
		Description:  input.Description,
		Flavor:       input.Flavor,
		Quantity:     input.Quantity,
		Status:       entity.StatusEnum(input.Status),
	}

	productUpdated, err := c.productRepository.Update(product)

	if err != nil {
		return nil, err
	}

	return &OutputUpdateProductDto{
		ID:           productUpdated.ID,
		Name:         productUpdated.Name,
		PriceInCents: productUpdated.PriceInCents,
		Description:  productUpdated.Description,
		Flavor:       productUpdated.Flavor,
		Quantity:     productUpdated.Quantity,
		CreatedAt:    productUpdated.CreatedAt,
		UpdatedAt:    productUpdated.UpdatedAt,
	}, nil
}
