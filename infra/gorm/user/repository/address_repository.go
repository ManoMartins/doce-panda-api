package repository

import (
	"doce-panda/domain/user/entity"
	"doce-panda/infra/gorm/user/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type AddressRepositoryDb struct {
	Db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepositoryDb {
	return &AddressRepositoryDb{Db: db}
}

func (r AddressRepositoryDb) FindById(ID string) (*entity.Address, error) {
	var addressModel model.Address

	r.Db.First(&addressModel, "id = ?", ID)

	if addressModel.ID == "" {
		return nil, fmt.Errorf("O endereço não foi encontrado")
	}

	return entity.NewAddress(entity.Address{
		ID:           addressModel.ID,
		City:         addressModel.City,
		State:        addressModel.State,
		Street:       addressModel.Street,
		Number:       addressModel.Number,
		ZipCode:      addressModel.ZipCode,
		Neighborhood: addressModel.Neighborhood,
		IsMain:       addressModel.IsMain,
		UserID:       addressModel.UserID,
		CreatedAt:    addressModel.CreatedAt,
		UpdatedAt:    addressModel.UpdatedAt,
	})
}

func (r AddressRepositoryDb) FindByMain() (*entity.Address, error) {
	var addressModel model.Address

	r.Db.First(&addressModel, "is_main = ?", true)

	if addressModel.ID == "" {
		return nil, fmt.Errorf("O endereço não foi encontrado")
	}

	return entity.NewAddress(entity.Address{
		ID:           addressModel.ID,
		City:         addressModel.City,
		State:        addressModel.State,
		Street:       addressModel.Street,
		Number:       addressModel.Number,
		ZipCode:      addressModel.ZipCode,
		Neighborhood: addressModel.Neighborhood,
		IsMain:       addressModel.IsMain,
		UserID:       addressModel.UserID,
		CreatedAt:    addressModel.CreatedAt,
		UpdatedAt:    addressModel.UpdatedAt,
	})
}

func (r AddressRepositoryDb) Delete(ID string) error {
	address := entity.Address{ID: ID}

	err := r.Db.Delete(&address).Error

	if err != nil {
		return err
	}

	return nil
}

func (r AddressRepositoryDb) FindAllByUserId(userID string) (*[]entity.Address, error) {
	var addressesModel []model.Address

	err := r.Db.Find(&addressesModel, "user_id = ?", userID).Error

	if err != nil {
		return nil, err
	}

	var addresses []entity.Address

	for _, addressModel := range addressesModel {
		address, err := entity.NewAddress(entity.Address{
			ID:           addressModel.ID,
			City:         addressModel.City,
			State:        addressModel.State,
			Street:       addressModel.Street,
			Number:       addressModel.Number,
			ZipCode:      addressModel.ZipCode,
			Neighborhood: addressModel.Neighborhood,
			IsMain:       addressModel.IsMain,
			UserID:       addressModel.UserID,
			CreatedAt:    addressModel.CreatedAt,
			UpdatedAt:    addressModel.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func (r AddressRepositoryDb) Update(address entity.Address) error {
	addressModel := model.Address{
		ID:           address.ID,
		City:         address.City,
		State:        address.State,
		Street:       address.Street,
		Number:       address.Number,
		ZipCode:      address.ZipCode,
		Neighborhood: address.Neighborhood,
		IsMain:       address.IsMain,
		UserID:       address.UserID,
		CreatedAt:    address.CreatedAt,
		UpdatedAt:    address.UpdatedAt,
	}

	err := r.Db.Save(&addressModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (r AddressRepositoryDb) Create(address entity.Address) error {
	addressModel := model.Address{
		ID:           address.ID,
		City:         address.City,
		State:        address.State,
		Street:       address.Street,
		Number:       address.Number,
		ZipCode:      address.ZipCode,
		Neighborhood: address.Neighborhood,
		IsMain:       address.IsMain,
		UserID:       address.UserID,
		CreatedAt:    address.CreatedAt,
		UpdatedAt:    address.UpdatedAt,
	}

	err := r.Db.Create(&addressModel).Error

	if err != nil {
		return err
	}

	return nil
}
