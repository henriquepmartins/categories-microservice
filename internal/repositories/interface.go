package repositories

import "github.com/henriquepermartins/categories-ms/internal/entities"

type CategoryRepositoryInterface interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
