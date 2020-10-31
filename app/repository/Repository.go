package repository

// Repository represents generic interface for interacting with DB
type Repository interface {
	Add(uow *UnitOfWork, target interface{}) error
	Delete(uow *UnitOfWork, targetID interface{}, target interface{}, where ...interface{}) error
	Update(uow *UnitOfWork, targetID interface{}, target interface{}) error
}

// Repo implement Repository
type Repo struct{}

// NewRepo returns a new repository object
func NewRepo() *Repo {
	return &Repo{}
}

// Add specified Entity
func (repository *Repo) Add(uow *UnitOfWork, entity interface{}) error {
	if err := uow.DB.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

// Update specified Entity
func (repository *Repo) Update(uow *UnitOfWork, entity interface{}) error {
	if err := uow.DB.Model(entity).Update(entity).Error; err != nil {
		return err
	}
	return nil
}

// Delete specified Entity
func (repository *Repo) Delete(uow *UnitOfWork, entity interface{}, where ...interface{}) error {
	if err := uow.DB.Delete(entity, where...).Error; err != nil {
		return err
	}
	return nil
}
