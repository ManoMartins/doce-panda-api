package user

import (
	"doce-panda/domain/user/repository"
	"doce-panda/usecase/user/dtos"
)

type FindAllAddressUseCase struct {
	addressRepo repository.AddressRepositoryInterface
}

func NewFindAllAddressUseCase(addressRepo repository.AddressRepositoryInterface) *FindAllAddressUseCase {
	return &FindAllAddressUseCase{
		addressRepo: addressRepo,
	}
}

func (c FindAllAddressUseCase) Execute(input dtos.InputFindAllAddressDto) (*[]dtos.OutputFindAllAddressDto, error) {
	addresses, err := c.addressRepo.FindAllByUserId(input.UserID)

	if err != nil {
		return nil, err
	}

	var output []dtos.OutputFindAllAddressDto

	for _, address := range *addresses {
		output = append(output, dtos.OutputFindAllAddressDto{
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
		})
	}

	return &output, nil
}
