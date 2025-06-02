package repositories

import "github.com/henriquepermartins/categories-ms/internal/entities"

type categoryRepository struct {
	db []*entities.Category
}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *categoryRepository) Save(category *entities.Category) error {
	r.db = append(r.db, category)

	return nil
}

func (r *categoryRepository) List() ([]*entities.Category, error) {
	return r.db, nil
}
