package user

import (
	"doce-panda/domain/user/repository"
	"doce-panda/usecase/user/dtos"
	"fmt"
)

type DestroyAddressUseCase struct {
	AddressRepository repository.AddressRepositoryInterface
}

func NewDestroyAddressUseCase(productRepository repository.AddressRepositoryInterface) *DestroyAddressUseCase {
	return &DestroyAddressUseCase{AddressRepository: productRepository}
}

func (c DestroyAddressUseCase) Execute(input dtos.InputDestroyAddressDto) error {
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
