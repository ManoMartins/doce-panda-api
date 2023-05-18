package product

import (
	"doce-panda/businessController/product/dtos"
	"doce-panda/domain/product/repository"
)

type DisableProductBusinessController struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewDisableProductBusinessController(productRepository repository.ProductRepositoryInterface) *DisableProductBusinessController {
	return &DisableProductBusinessController{ProductRepository: productRepository}
}

func (c DisableProductBusinessController) Execute(input dtos.InputDisableProductDto) error {
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
