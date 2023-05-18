package user

import (
	"doce-panda/businessController/user/dtos"
	"doce-panda/domain/user/entity"
	"doce-panda/domain/user/repository"
	"time"
)

type UpdateAddressBusinessController struct {
	addressRepository repository.AddressRepositoryInterface
}

func NewUpdateAddressBusinessController(addressRepository repository.AddressRepositoryInterface) *UpdateAddressBusinessController {
	return &UpdateAddressBusinessController{
		addressRepository: addressRepository,
	}
}

func (c UpdateAddressBusinessController) Execute(input dtos.InputUpdateAddressDto) (*dtos.OutputUpdateAddressDto, error) {
	addressFound, err := c.addressRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	err = addressFound.Update(entity.AddressUpdateProps{
		City:         input.City,
		State:        input.State,
		Street:       input.Street,
		Number:       input.Number,
		ZipCode:      input.ZipCode,
		Neighborhood: input.Neighborhood,
		IsMain:       input.IsMain,
	})

	addressFound.UpdatedAt = time.Now()

	if err != nil {
		return nil, err
	}

	err = c.addressRepository.Update(*addressFound)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputUpdateAddressDto{
		ID:           addressFound.ID,
		City:         addressFound.City,
		State:        addressFound.State,
		Street:       addressFound.Street,
		Number:       addressFound.Number,
		ZipCode:      addressFound.ZipCode,
		Neighborhood: addressFound.Neighborhood,
		IsMain:       addressFound.IsMain,
		CreatedAt:    addressFound.CreatedAt,
		UpdatedAt:    addressFound.UpdatedAt,
	}, nil
}
