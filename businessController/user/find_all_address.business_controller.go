package user

import (
	"doce-panda/businessController/user/dtos"
	"doce-panda/domain/user/repository"
)

type FindAllAddressBusinessController struct {
	addressRepo repository.AddressRepositoryInterface
}

func NewFindAllAddressBusinessController(addressRepo repository.AddressRepositoryInterface) *FindAllAddressBusinessController {
	return &FindAllAddressBusinessController{
		addressRepo: addressRepo,
	}
}

func (c FindAllAddressBusinessController) Execute(input dtos.InputFindAllAddressDto) (*[]dtos.OutputFindAllAddressDto, error) {
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
