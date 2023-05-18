package user

import (
	"doce-panda/businessController/user/dtos"
	"doce-panda/domain/user/entity"
	"doce-panda/domain/user/repository"
)

type CreateAddressBusinessController struct {
	addressRepository repository.AddressRepositoryInterface
}

func NewCreateAddressBusinessController(addressRepository repository.AddressRepositoryInterface) *CreateAddressBusinessController {
	return &CreateAddressBusinessController{addressRepository: addressRepository}
}

func (c CreateAddressBusinessController) Execute(input dtos.InputCreateAddressDto) (*dtos.OutputCreateAddressDto, error) {
	addressMain, _ := c.addressRepository.FindByMain()

	if addressMain != nil {
		addressMain.DisableMain()
		err := c.addressRepository.Update(*addressMain)

		if err != nil {
			return nil, err
		}
	}

	address, err := entity.NewAddress(entity.Address{
		City:         input.City,
		State:        input.State,
		Street:       input.Street,
		Number:       input.Number,
		ZipCode:      input.ZipCode,
		Neighborhood: input.Neighborhood,
		IsMain:       input.IsMain,
		UserID:       input.UserID,
	})

	if err != nil {
		return nil, err
	}

	err = c.addressRepository.Create(*address)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputCreateAddressDto{
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
	}, nil
}
