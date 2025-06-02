package useCase

import (
	"log"

	"github.com/henriquepermartins/categories-ms/internal/entities"
	"github.com/henriquepermartins/categories-ms/internal/repositories"
)

type createCategoryUseCase struct {
	categoryRepository repositories.CategoryRepositoryInterface
}

func NewCreateCategoryUseCase(repository repositories.CategoryRepositoryInterface) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (usecase *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	log.Println(category)
	err = usecase.categoryRepository.Save(category)

	if err != nil {
		return err
	}

	return nil
}
