package payment

import (
	"doce-panda/businessController/payment/dtos"
	"doce-panda/domain/payment/repository"
)

type FindByIdCreditCardBusinessController struct {
	creditCardRepository repository.CreditCardRepositoryInterface
}

func NewFindByIdCreditCardBusinessController(creditCardRepository repository.CreditCardRepositoryInterface) *FindByIdCreditCardBusinessController {
	return &FindByIdCreditCardBusinessController{creditCardRepository: creditCardRepository}
}

func (c FindByIdCreditCardBusinessController) Execute(input dtos.InputFindByIdCreditCardDto) (*dtos.OutputFindByIdCreditCardDto, error) {
	creditCard, err := c.creditCardRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputFindByIdCreditCardDto{
		ID:                 creditCard.ID,
		CardLastNumber:     creditCard.CardLastNumber,
		CardHolder:         creditCard.CardHolder,
		CardIdentification: creditCard.CardIdentification,
		CardSecurityCode:   creditCard.CardSecurityCode,
		CardExpirationDate: creditCard.CardExpirationDate,
		CardBrand:          creditCard.CardBrand,
		CreatedAt:          creditCard.CreatedAt,
		UpdatedAt:          creditCard.UpdatedAt,
	}, nil
}
