package model

import (
	"time"
)

type User struct {
	ID             string    `json:"id" gorm:"type:uuid;primary_key"`
	Name           string    `json:"name" gorm:"type:varchar(255)"`
	Gender         string    `json:"gender" gorm:"type:varchar(255)"`
	Password       string    `json:"password" gorm:"type:varchar(255)"`
	PhoneNumber    string    `json:"phone_number" gorm:"type:varchar(11)"`
	DocumentNumber string    `json:"document_number" gorm:"type:varchar(11)"`
	RewardPoints   int       `json:"reward_points" gorm:"type:integer"`
	Email          string    `json:"email" gorm:"type:varchar(255)"`
	Addresses      []Address `json:"addresses"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
