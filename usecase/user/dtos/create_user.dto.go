package dtos

import "time"

type InputCreateUserDto struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	Password       string `json:"password"`
	PhoneNumber    string `json:"phoneNumber"`
	DocumentNumber string `json:"documentNumber"`
}
type OutputCreateUserDto struct {
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
