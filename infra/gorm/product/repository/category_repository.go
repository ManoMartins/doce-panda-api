package repository

import (
	"doce-panda/domain/product/entity"
	"doce-panda/infra/gorm/product/model"
	"github.com/jinzhu/gorm"
	"time"
)

type CategoryRepositoryDb struct {
	Db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryDb {
	return &CategoryRepositoryDb{Db: db}
}

func (c CategoryRepositoryDb) Delete(ID string) error {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepositoryDb) FindAll() (*[]entity.Category, error) {
	var categoriesModel []model.Category

	err := c.Db.Find(&categoriesModel).Error

	if err != nil {
		return nil, err
	}

	var categories []entity.Category

	for _, categoryModel := range categoriesModel {
		category, err := entity.NewCategory(entity.Category{
			ID:          categoryModel.ID,
			Description: categoryModel.Description,
			CreatedAt:   categoryModel.CreatedAt,
			UpdatedAt:   categoryModel.UpdatedAt,
		})

		if err != nil {
			return nil, err
		}

		categories = append(categories, *category)
	}

	return &categories, nil
}

func (c CategoryRepositoryDb) Update(category entity.Category) error {
	categoryModel := model.Category{
		ID:          category.ID,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}

	err := c.Db.Save(&categoryModel).Error

	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepositoryDb) Create(category entity.Category) error {
	categoryModel := model.Category{
		ID:          category.ID,
		Description: category.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := c.Db.Create(&categoryModel).Error

	if err != nil {
		return err
	}

	return nil
}
