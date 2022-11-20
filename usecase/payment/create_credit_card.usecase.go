package payment

import (
	"doce-panda/domain/payment/entity"
	"doce-panda/domain/payment/repository"
	"doce-panda/usecase/payment/dtos"
)

type CreateCreditCardUseCase struct {
	creditCardRepository repository.CreditCardRepositoryInterface
}

func NewCreateCreditCardUseCase(creditCardRepository repository.CreditCardRepositoryInterface) *CreateCreditCardUseCase {
	return &CreateCreditCardUseCase{creditCardRepository: creditCardRepository}
}

func (c CreateCreditCardUseCase) Execute(input dtos.InputCreateCreditCardDto) (*dtos.OutputCreateCreditCardDto, error) {
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
