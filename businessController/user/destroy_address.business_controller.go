package user

import (
	"doce-panda/businessController/user/dtos"
	"doce-panda/domain/user/repository"
	"fmt"
)

type DestroyAddressBusinessController struct {
	AddressRepository repository.AddressRepositoryInterface
}

func NewDestroyAddressBusinessController(productRepository repository.AddressRepositoryInterface) *DestroyAddressBusinessController {
	return &DestroyAddressBusinessController{AddressRepository: productRepository}
}

func (c DestroyAddressBusinessController) Execute(input dtos.InputDestroyAddressDto) error {
	addressFound, err := c.AddressRepository.FindById(input.AddressID)

	if err != nil {
		return err
	}

	if addressFound.UserID != input.ID {
		return fmt.Errorf("Endereço não pertece ao usuário")
	}

	err = c.AddressRepository.Delete(input.AddressID)

	if err != nil {
		return err
	}

	return nil
}
