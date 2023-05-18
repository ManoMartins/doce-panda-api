package dtos

import (
	"doce-panda/domain/user/entity"
	"time"
)

type InputFindUserDto struct {
	ID string
}

type OutputFindUserDto struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Gender         string           `json:"gender"`
	PhoneNumber    string           `json:"phoneNumber"`
	DocumentNumber string           `json:"documentNumber"`
	RewardPoints   int              `json:"rewardPoints"`
	Email          string           `json:"email"`
	Addresses      []entity.Address `json:"addresses"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
}
