package entity

import (
	"doce-panda/domain/product/entity"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID             string    `json:"id" validate:"required"`
	Name           string    `json:"name" validate:"required"`
	Gender         string    `json:"gender" validate:"required,oneof='male' 'female'"`
	Password       string    `json:"password"`
	PhoneNumber    string    `json:"phone_number" validate:"required"`
	DocumentNumber string    `json:"document_number" validate:"required"`
	RewardPoints   int       `json:"reward_points" validate:"numeric"`
	Email          string    `json:"email" validate:"required,email"`
	Addresses      []Address `json:"addresses"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserUpdateProps struct {
	Name           string
	Email          string
	Gender         string
	PhoneNumber    string
	DocumentNumber string
}

type UserInterface interface {
	AddRewardPointsBuyProduct(product []entity.Product)
	Validate(props User) error
	Update(userUpdate UserUpdateProps) error
}

func NewUser(user User) (*User, error) {
	u := User{
		ID:             user.ID,
		Name:           user.Name,
		Gender:         user.Gender,
		Password:       user.Password,
		PhoneNumber:    user.PhoneNumber,
		DocumentNumber: user.DocumentNumber,
		RewardPoints:   user.RewardPoints,
		Email:          user.Email,
		Addresses:      user.Addresses,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}

	if user.ID == "" {
		u.ID = uuid.NewString()
	}

	err := u.Validate(u)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (u *User) Validate(props User) error {
	var Validator = validator.New()

	return Validator.Struct(props)
}

func (u *User) Update(userUpdate UserUpdateProps) error {
	u.Name = userUpdate.Name
	u.Gender = userUpdate.Gender
	u.PhoneNumber = userUpdate.PhoneNumber
	u.DocumentNumber = userUpdate.DocumentNumber
	u.Email = userUpdate.Email

	err := u.Validate(*u)

	if err != nil {
		return err
	}

	return nil

}

func (u *User) AddRewardPointsBuyProduct(products []entity.Product) {
	totalRewardPoints := 0

	for _, p := range products {
		totalRewardPoints += p.PriceInCents * p.Quantity
	}

	u.RewardPoints = totalRewardPoints
}
