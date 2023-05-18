package dtos

import "time"

type InputCreateCreditCardDto struct {
	CardLastNumber     string `json:"cardLastNumber"`
	CardHolder         string `json:"cardHolder"`
	CardIdentification string `json:"cardIdentification"`
	CardSecurityCode   string `json:"cardSecurityCode"`
	CardExpirationDate string `json:"cardExpirationDate"`
	CardBrand          string `json:"cardBrand"`
}

type OutputCreateCreditCardDto struct {
	ID                 string    `json:"id"`
	CardLastNumber     string    `json:"cardLastNumber"`
	CardHolder         string    `json:"cardHolder"`
	CardIdentification string    `json:"cardIdentification"`
	CardSecurityCode   string    `json:"cardSecurityCode"`
	CardExpirationDate string    `json:"cardExpirationDate"`
	CardBrand          string    `json:"cardBrand"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}
