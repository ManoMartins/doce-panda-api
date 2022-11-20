package entity

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type CreditCard struct {
	ID                 string    `json:"id" validate:"required"`
	CardLastNumber     string    `json:"cardLastNumber" validate:"required"`
	CardHolder         string    `json:"cardHolder" validate:"required"`
	CardIdentification string    `json:"cardIdentification" validate:"required"`
	CardSecurityCode   string    `json:"cardSecurityCode" validate:"required"`
	CardExpirationDate string    `json:"cardExpirationDate" validate:"required"`
	CardBrand          string    `json:"cardBrand" validate:"required"`
	TotalInCents       int       `json:"totalInCents"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

type ProductInterface interface {
	Validate() error
}

func NewCreditCard(creditCard CreditCard) (*CreditCard, error) {
	p := CreditCard{
		ID:                 creditCard.ID,
		CardLastNumber:     creditCard.CardLastNumber,
		CardHolder:         creditCard.CardHolder,
		CardIdentification: creditCard.CardIdentification,
		CardSecurityCode:   creditCard.CardSecurityCode,
		CardExpirationDate: creditCard.CardExpirationDate,
		CardBrand:          creditCard.CardBrand,
		CreatedAt:          creditCard.CreatedAt,
		UpdatedAt:          creditCard.UpdatedAt,
	}

	if creditCard.ID == "" {
		p.ID = uuid.NewString()
	}

	err := p.Validate()

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (c *CreditCard) Validate() error {
	var Validator = validator.New()

	return Validator.Struct(c)
}
