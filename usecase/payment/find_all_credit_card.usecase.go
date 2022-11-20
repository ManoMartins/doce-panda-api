package payment

import (
	"doce-panda/domain/payment/repository"
	"doce-panda/usecase/payment/dtos"
)

type FindAllCreditCardUseCase struct {
	creditCardRepository repository.CreditCardRepositoryInterface
}

func NewFindAllCreditCardUseCase(creditCardRepository repository.CreditCardRepositoryInterface) *FindAllCreditCardUseCase {
	return &FindAllCreditCardUseCase{creditCardRepository: creditCardRepository}
}

func (c FindAllCreditCardUseCase) Execute() (*[]dtos.OutputFindAllCreditCardDto, error) {
	creditCards, err := c.creditCardRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var output []dtos.OutputFindAllCreditCardDto
	for _, creditCard := range *creditCards {
		output = append(output, dtos.OutputFindAllCreditCardDto{
			ID:                 creditCard.ID,
			CardLastNumber:     creditCard.CardLastNumber,
			CardHolder:         creditCard.CardHolder,
			CardIdentification: creditCard.CardIdentification,
			CardSecurityCode:   creditCard.CardSecurityCode,
			CardExpirationDate: creditCard.CardExpirationDate,
			CardBrand:          creditCard.CardBrand,
			CreatedAt:          creditCard.CreatedAt,
			UpdatedAt:          creditCard.UpdatedAt,
		})
	}

	return &output, nil
}
