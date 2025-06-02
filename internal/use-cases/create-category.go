package useCase

import (
	"log"

	"github.com/henriquepermartins/categories-ms/internal/entities"
)

type createCategoryUseCase struct {
	// will receive db instance
}

func NewCreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}
}

func (usecase *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// TODO: persist entity in db
	log.Println(category)

	return nil
}
