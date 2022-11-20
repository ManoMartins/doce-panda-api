package repository

import (
	"doce-panda/domain/payment/entity"
	"doce-panda/infra/gorm/payment/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type CreditCardRepositoryDb struct {
	Db *gorm.DB
}

func NewCreditCardRepository(db *gorm.DB) *CreditCardRepositoryDb {
	return &CreditCardRepositoryDb{Db: db}
}

func (c CreditCardRepositoryDb) FindById(ID string) (*entity.CreditCard, error) {
	var creditCardModel model.CreditCard

	c.Db.First(&creditCardModel, "id = ?", ID)

	if creditCardModel.ID == "" {
		return nil, fmt.Errorf("Cartão não foi encontrado")
	}

	return entity.NewCreditCard(entity.CreditCard{
		ID:                 creditCardModel.ID,
		CardLastNumber:     creditCardModel.CardLastNumber,
		CardHolder:         creditCardModel.CardHolder,
		CardIdentification: creditCardModel.CardIdentification,
		CardSecurityCode:   creditCardModel.CardSecurityCode,
		CardExpirationDate: fmt.Sprintf("%s/%s", creditCardModel.CardMonth, creditCardModel.CardYear),
		CardBrand:          creditCardModel.CardBrand,
		CreatedAt:          creditCardModel.CreatedAt,
		UpdatedAt:          creditCardModel.UpdatedAt,
	})
}

func (c CreditCardRepositoryDb) FindAll() (*[]entity.CreditCard, error) {
	var creditCardsModel []model.CreditCard

	err := c.Db.Find(&creditCardsModel).Error

	if err != nil {
		return nil, err
	}

	var creditCards []entity.CreditCard

	for _, creditCardModel := range creditCardsModel {
		creditCard, err := entity.NewCreditCard(entity.CreditCard{
			ID:                 creditCardModel.ID,
			CardLastNumber:     creditCardModel.CardLastNumber,
			CardHolder:         creditCardModel.CardHolder,
			CardIdentification: creditCardModel.CardIdentification,
			CardSecurityCode:   creditCardModel.CardSecurityCode,
			CardExpirationDate: fmt.Sprintf("%d/%d", creditCardModel.CardMonth, creditCardModel.CardYear),
			CardBrand:          creditCardModel.CardBrand,
			CreatedAt:          creditCardModel.CreatedAt,
			UpdatedAt:          creditCardModel.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		creditCards = append(creditCards, *creditCard)
	}

	return &creditCards, nil
}

func (c CreditCardRepositoryDb) Create(creditCard entity.CreditCard) error {
	cardExpirationDate := strings.Split(creditCard.CardExpirationDate, "/")

	cardMonth := cardExpirationDate[0]
	CardYear := cardExpirationDate[1]

	creditCardModel := model.CreditCard{
		ID:                 creditCard.ID,
		CardLastNumber:     creditCard.CardLastNumber,
		CardHolder:         creditCard.CardHolder,
		CardIdentification: creditCard.CardIdentification,
		CardSecurityCode:   creditCard.CardSecurityCode,
		CardMonth:          cardMonth,
		CardYear:           CardYear,
		CardBrand:          creditCard.CardBrand,
		CreatedAt:          creditCard.CreatedAt,
		UpdatedAt:          creditCard.UpdatedAt,
	}

	err := c.Db.Create(&creditCardModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (c CreditCardRepositoryDb) Delete(ID string) error {
	creditCard := entity.CreditCard{ID: ID}

	err := c.Db.Delete(&creditCard).Error

	if err != nil {
		return err
	}

	return nil
}
