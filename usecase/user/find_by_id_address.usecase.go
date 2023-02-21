package user

import (
	"doce-panda/domain/user/repository"
	"doce-panda/usecase/user/dtos"
)

type FindByIdAddressUseCase struct {
	addressRepo repository.AddressRepositoryInterface
}

func NewFindByIdAddressUseCase(addressRepo repository.AddressRepositoryInterface) *FindByIdAddressUseCase {
	return &FindByIdAddressUseCase{
		addressRepo: addressRepo,
	}
}

func (c FindByIdAddressUseCase) Execute(input dtos.InputFindByIdAddressDto) (*dtos.OutputFindByIdAddressDto, error) {
	address, err := c.addressRepo.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	output := dtos.OutputFindByIdAddressDto{
		ID:           address.ID,
		City:         address.City,
		State:        address.State,
		Street:       address.Street,
		Number:       address.Number,
		ZipCode:      address.ZipCode,
		Neighborhood: address.Neighborhood,
		IsMain:       address.IsMain,
		UserID:       address.UserID,
		CreatedAt:    address.CreatedAt,
		UpdatedAt:    address.UpdatedAt,
	}

	return &output, nil
}
