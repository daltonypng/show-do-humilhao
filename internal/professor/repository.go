package professor

import (
	"daltonypng/show-do-humilhao/internal/covenant"
	"daltonypng/show-do-humilhao/internal/entity"
)

type Repository struct {
	database covenant.Database
}

func NewRepository(database covenant.Database) *Repository {
	return &Repository{
		database: database,
	}
}

func (repository *Repository) FindByCredentials(email string, password string) error {

	professor := &entity.Professor{}

	return repository.database.Read(&professor, "email = ? and password = ?", email, password)

}

func (repository *Repository) FindByEmail(email string) (*entity.Professor, error) {

	professor := &entity.Professor{}

	err := repository.database.Read(&professor, "email = ? ", []string{email})

	return professor, err

}

func (repository *Repository) Create(professor *entity.Professor) error {

	return repository.database.Create(professor)
}
