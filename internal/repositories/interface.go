package repositories

type CategoryRepositoryInterface interface {
	Save(name string) error
}