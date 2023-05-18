package dtos

import "time"

type OutputFindAllUserDto struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Gender         string    `json:"gender"`
	PhoneNumber    string    `json:"phoneNumber"`
	DocumentNumber string    `json:"documentNumber"`
	RewardPoints   int       `json:"rewardPoints"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
