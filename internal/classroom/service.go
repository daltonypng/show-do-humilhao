package classroom

import (
	"daltonypng/show-do-humilhao/internal/apperror"
	"daltonypng/show-do-humilhao/internal/entity"
	"errors"
	"math/rand/v2"
)

const (
	min uint = 100000
	max uint = 999999
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) Create(classroom *entity.Classroom) error {

	if len(classroom.Name) <= 0 {
		return errors.New(apperror.ClassroomEmptyName)

	} else if classroom.ProfessorID <= 0 {
		return errors.New(apperror.ClassroomInvalidProfessor)

	}

	classroom.ID = generateNewClassroomID()

	return service.repository.Create(classroom)
}

func (service *Service) FindByID(ID uint) (*entity.Classroom, error) {
	return nil, nil
}

func generateNewClassroomID() uint {

	return min + rand.UintN(max-min+1)
}
