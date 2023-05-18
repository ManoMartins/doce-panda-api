package product

import (
	"doce-panda/businessController/product/dtos"
	"doce-panda/domain/product/repository"
)

type EnableProductBusinessController struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewEnableProductBusinessController(productRepository repository.ProductRepositoryInterface) *EnableProductBusinessController {
	return &EnableProductBusinessController{ProductRepository: productRepository}
}

func (c EnableProductBusinessController) Execute(input dtos.InputEnableProductDto) error {
	product, err := c.ProductRepository.FindById(input.ID)

	err = product.Enable()

	if err != nil {
		return err
	}

	err = c.ProductRepository.Update(*product)

	if err != nil {
		return err
	}

	return nil
}
