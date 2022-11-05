package user

import (
	"doce-panda/domain/user/entity"
	"doce-panda/domain/user/repository"
	"doce-panda/usecase/user/dtos"
	"time"
)

type UpdateAddressUseCase struct {
	addressRepository repository.AddressRepositoryInterface
}

func NewUpdateAddressUseCase(addressRepository repository.AddressRepositoryInterface) *UpdateAddressUseCase {
	return &UpdateAddressUseCase{
		addressRepository: addressRepository,
	}
}

func (c UpdateAddressUseCase) Execute(input dtos.InputUpdateAddressDto) (*dtos.OutputUpdateAddressDto, error) {
	addressFound, err := c.addressRepository.Find(input.ID)

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
