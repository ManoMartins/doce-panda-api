package payment

import (
	"doce-panda/domain/payment/repository"
	"doce-panda/usecase/payment/dtos"
)

type DeleteCreditCardUseCase struct {
	CreditCardRepository repository.CreditCardRepositoryInterface
}

func NewDeleteCreditCardUseCase(creditCardRepository repository.CreditCardRepositoryInterface) *DeleteCreditCardUseCase {
	return &DeleteCreditCardUseCase{CreditCardRepository: creditCardRepository}
}

func (c DeleteCreditCardUseCase) Execute(input dtos.InputDeleteCreditCardDto) error {
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
