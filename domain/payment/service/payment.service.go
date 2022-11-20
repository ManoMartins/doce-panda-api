package service

import paymentEntity "doce-panda/domain/payment/entity"

func Total(creditCards []paymentEntity.CreditCard) int {
	totalInCents := 0

	for _, creditCard := range creditCards {
		totalInCents += creditCard.TotalInCents
	}

	return totalInCents
}
