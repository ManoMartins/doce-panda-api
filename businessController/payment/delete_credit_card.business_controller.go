package payment

import (
	"doce-panda/businessController/payment/dtos"
	"doce-panda/domain/payment/repository"
)

type DeleteCreditCardBusinessController struct {
	CreditCardRepository repository.CreditCardRepositoryInterface
}

func NewDeleteCreditCardBusinessController(creditCardRepository repository.CreditCardRepositoryInterface) *DeleteCreditCardBusinessController {
	return &DeleteCreditCardBusinessController{CreditCardRepository: creditCardRepository}
}

func (c DeleteCreditCardBusinessController) Execute(input dtos.InputDeleteCreditCardDto) error {
	_, err := c.CreditCardRepository.FindById(input.ID)

	if err != nil {
		return err
	}

	err = c.CreditCardRepository.Delete(input.ID)

	if err != nil {
		return err
	}

	return nil
}
