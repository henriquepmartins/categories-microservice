package useCase

import (
	"github.com/henriquepermartins/categories-ms/internal/entities"
	"github.com/henriquepermartins/categories-ms/internal/repositories"
)

type listCategoryUseCase struct {
	categoryRepository repositories.CategoryRepositoryInterface
}

func NewListCategoryUseCase(repository repositories.CategoryRepositoryInterface) *listCategoryUseCase {
	return &listCategoryUseCase{repository}
}

func (usecase *listCategoryUseCase) Execute() ([]*entities.Category, error) {
	categories, err := usecase.categoryRepository.List()

	if err != nil {
		return nil, err
	}

	return categories, nil
}
