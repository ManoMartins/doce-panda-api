package product

import (
	"doce-panda/domain/product/repository"
	"doce-panda/usecase/product/dtos"
)

type DisableProductUseCase struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewDisableProductUseCase(productRepository repository.ProductRepositoryInterface) *DisableProductUseCase {
	return &DisableProductUseCase{ProductRepository: productRepository}
}

func (c DisableProductUseCase) Execute(input dtos.InputDisableProductDto) error {
	product, err := c.ProductRepository.FindById(input.ID)

	err = product.Disable()

	if err != nil {
		return err
	}

	err = c.ProductRepository.Update(*product)

	if err != nil {
		return err
	}

	return nil
}
