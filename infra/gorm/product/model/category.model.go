package model

import "time"

type Category struct {
	ID          string    `json:"id" gorm:"type:uuid;primary_key"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
