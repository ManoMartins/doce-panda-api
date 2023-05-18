package product

import (
	"doce-panda/businessController/product/dtos"
	"doce-panda/domain/product/entity"
	"doce-panda/domain/product/repository"
)

type CreateCategoryBusinessController struct {
	categoryRepository repository.CategoryRepositoryInterface
}

func NewCreateCategoryBusinessController(categoryRepository repository.CategoryRepositoryInterface) *CreateCategoryBusinessController {
	return &CreateCategoryBusinessController{categoryRepository: categoryRepository}
}

func (c CreateCategoryBusinessController) Execute(input dtos.InputCreateCategoryDto) (*dtos.OutputCreateCategoryDto, error) {
	category, err := entity.NewCategory(entity.Category{
		Description: input.Description,
	})

	if err != nil {
		return nil, err
	}

	err = c.categoryRepository.Create(*category)

	if err != nil {
		return nil, err
	}

	return &dtos.OutputCreateCategoryDto{
		ID:          category.ID,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}
