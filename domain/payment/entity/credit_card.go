package entity

type CreditCard struct {
	ID                 string `json:"id" validate:"required"`
	CardLastNumber     string `json:"cardLastNumber" validate:"required"`
	CardHolder         string `json:"cardHolder" validate:"required"`
	CardIdentification string `json:"cardIdentification" validate:"required"`
	CardSecurityCode   string `json:"cardSecurityCode" validate:"required"`
	CardExpirationDate string `json:"cardExpirationDate" validate:"required"`
	CardBrand          string `json:"cardBrand" validate:"required"`
}
