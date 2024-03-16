package classroom

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

func (repository *Repository) Create(classroom *entity.Classroom) error {
	return repository.database.Create(classroom)
}

func (repository *Repository) FindByID(ID uint) (*entity.Classroom, error) {
	classroom := &entity.Classroom{}
	err := repository.database.Read(&classroom, ID)

	if err != nil {
		return nil, err
	}

	return classroom, nil
}

func (repository *Repository) Delete(classroom *entity.Classroom) error {
	return repository.database.Delete(classroom)
}
