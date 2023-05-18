package payment

import (
	"doce-panda/businessController/payment/dtos"
	"doce-panda/domain/payment/entity"
	"doce-panda/domain/payment/repository"
)

type CreateCreditCardBusinessController struct {
	creditCardRepository repository.CreditCardRepositoryInterface
}

func NewCreateCreditCardBusinessController(creditCardRepository repository.CreditCardRepositoryInterface) *CreateCreditCardBusinessController {
	return &CreateCreditCardBusinessController{creditCardRepository: creditCardRepository}
}

func (c CreateCreditCardBusinessController) Execute(input dtos.InputCreateCreditCardDto) (*dtos.OutputCreateCreditCardDto, error) {
	creditCard, err := entity.NewCreditCard(entity.CreditCard{
		CardLastNumber:     input.CardLastNumber,
		CardHolder:         input.CardHolder,
		CardIdentification: input.CardIdentification,
		CardSecurityCode:   input.CardSecurityCode,
		CardExpirationDate: input.CardExpirationDate,
		CardBrand:          input.CardBrand,
	})

	if err != nil {
		return nil, err
	}

	c.creditCardRepository.Create(*creditCard)

	return &dtos.OutputCreateCreditCardDto{
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
