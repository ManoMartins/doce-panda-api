package user

import (
	"doce-panda/businessController/user/dtos"
	"doce-panda/domain/user/repository"
)

type FindByIdAddressBusinessController struct {
	addressRepo repository.AddressRepositoryInterface
}

func NewFindByIdAddressBusinessController(addressRepo repository.AddressRepositoryInterface) *FindByIdAddressBusinessController {
	return &FindByIdAddressBusinessController{
		addressRepo: addressRepo,
	}
}

func (c FindByIdAddressBusinessController) Execute(input dtos.InputFindByIdAddressDto) (*dtos.OutputFindByIdAddressDto, error) {
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
