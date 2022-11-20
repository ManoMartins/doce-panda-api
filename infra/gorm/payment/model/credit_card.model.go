package model

import "time"

type CreditCard struct {
	ID                 string    `json:"id" gorm:"type:uuid;primary_key"`
	CardLastNumber     string    `json:"cardLastNumber" gorm:"type:varchar(255)"`
	CardHolder         string    `json:"cardHolder" gorm:"type:varchar(255)"`
	CardIdentification string    `json:"cardIdentification" gorm:"type:varchar(255)"`
	CardSecurityCode   string    `json:"cardSecurityCode" gorm:"type:varchar(3)"`
	CardMonth          string    `json:"cardMonth" gorm:"type:varchar(2)"`
	CardYear           string    `json:"cardYear" gorm:"type:varchar(4)"`
	CardBrand          string    `json:"cardBrand" gorm:"type:varchar(255)"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}
