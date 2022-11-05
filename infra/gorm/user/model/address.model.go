package model

import (
	"time"
)

type Address struct {
	ID           string    `json:"id" gorm:"type:uuid;primary_key"`
	City         string    `json:"city" gorm:"type:varchar(255)"`
	State        string    `json:"state" gorm:"type:varchar(255)"`
	Street       string    `json:"street" gorm:"type:varchar(255)"`
	Number       string    `json:"number" gorm:"type:varchar(255)"`
	ZipCode      string    `json:"zipCode" gorm:"type:varchar(255)"`
	Neighborhood string    `json:"neighborhood" gorm:"type:varchar(255)"`
	IsMain       bool      `json:"isMain" gorm:"type:bool"`
	User         User      `json:"user" gorm:"ForeignKey:UserID"`
	UserID       string    `json:"userId" gorm:"column:user_id;type:uuid;notnull"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
