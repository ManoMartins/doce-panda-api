package payment

import (
	"doce-panda/businessController/payment/dtos"
	"doce-panda/domain/payment/repository"
)

type FindAllCreditCardBusinessController struct {
	creditCardRepository repository.CreditCardRepositoryInterface
}

func NewFindAllCreditCardBusinessController(creditCardRepository repository.CreditCardRepositoryInterface) *FindAllCreditCardBusinessController {
	return &FindAllCreditCardBusinessController{creditCardRepository: creditCardRepository}
}

func (c FindAllCreditCardBusinessController) Execute() (*[]dtos.OutputFindAllCreditCardDto, error) {
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
