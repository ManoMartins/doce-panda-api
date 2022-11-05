package entity

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type Address struct {
	ID           string    `json:"id" validate:"required"`
	City         string    `json:"city" validate:"required"`
	State        string    `json:"state" validate:"required"`
	Street       string    `json:"street" validate:"required"`
	Number       string    `json:"number" validate:"required"`
	ZipCode      string    `json:"zipCode" validate:"required"`
	Neighborhood string    `json:"neighborhood" validate:"required"`
	IsMain       bool      `json:"isMain"`
	UserID       string    `json:"userId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type AddressUpdateProps struct {
	City         string `json:"city"`
	State        string `json:"state"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	ZipCode      string `json:"zipCode"`
	Neighborhood string `json:"neighborhood"`
	IsMain       bool   `json:"isMain"`
}

type AddressInterface interface {
	Validate(props Address) error
	Update(addressUpdate AddressUpdateProps) error
	DisableMain()
}

func NewAddress(address Address) (*Address, error) {
	a := Address{
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

	if address.ID == "" {
		a.ID = uuid.NewString()
	}

	err := a.Validate(a)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *Address) Validate(props Address) error {
	var Validator = validator.New()

	return Validator.Struct(props)
}

func (a *Address) Update(addressUpdate AddressUpdateProps) error {
	a.City = addressUpdate.City
	a.State = addressUpdate.State
	a.Street = addressUpdate.Street
	a.Number = addressUpdate.Number
	a.ZipCode = addressUpdate.ZipCode
	a.Neighborhood = addressUpdate.Neighborhood
	a.IsMain = addressUpdate.IsMain

	err := a.Validate(*a)

	if err != nil {
		return err
	}

	return nil
}

func (a *Address) DisableMain() {
	a.IsMain = false
}
