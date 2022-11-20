package payment

import (
	"doce-panda/domain/payment/repository"
	"doce-panda/usecase/payment/dtos"
)

type FindByIdCreditCardUseCase struct {
	creditCardRepository repository.CreditCardRepositoryInterface
}

func NewFindByIdCreditCardUseCase(creditCardRepository repository.CreditCardRepositoryInterface) *FindByIdCreditCardUseCase {
	return &FindByIdCreditCardUseCase{creditCardRepository: creditCardRepository}
}

func (c FindByIdCreditCardUseCase) Execute(input dtos.InputFindByIdCreditCardDto) (*dtos.OutputFindByIdCreditCardDto, error) {
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
