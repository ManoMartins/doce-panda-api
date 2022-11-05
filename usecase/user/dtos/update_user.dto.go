package dtos

import "time"

type InputUpdateUserDto struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	PhoneNumber    string `json:"phoneNumber"`
	DocumentNumber string `json:"documentNumber"`
}

type OutputUpdateUserDto struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Gender         string    `json:"gender"`
	PhoneNumber    string    `json:"phoneNumber"`
	DocumentNumber string    `json:"documentNumber"`
	RewardPoints   int       `json:"rewardPoints"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
