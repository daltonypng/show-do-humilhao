package classroom

import (
	"daltonypng/show-do-humilhao/internal/apperror"
	"daltonypng/show-do-humilhao/internal/entity"
	"errors"
	"math/rand/v2"
)

const (
	StatusCreating int = 0
	StatusPlaying  int = 1
	StatusFinished int = 2
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
	classroom.Status = StatusCreating

	return service.repository.Create(classroom)
}

func (service *Service) FindByID(ID uint) (*entity.Classroom, error) {

	if ID <= 0 {
		return nil, errors.New(apperror.ClassroomEmptyID)

	}

	classroom, err := service.repository.FindByID(ID)

	if err != nil {
		return nil, err

	} else if classroom.ID <= 0 {
		return nil, errors.New(apperror.ClassroomNotFound)

	}

	return classroom, nil
}

func (service *Service) UpdateStatusByID(ID uint, status int) error {

	if ID <= 0 {
		return errors.New(apperror.ClassroomEmptyID)

	} else if status < StatusCreating || status > StatusFinished {
		return errors.New(apperror.ClassroomInvalidStatus)

	}

	classroom, err := service.repository.FindByID(ID)

	if err != nil {
		return err
	}

	classroom.Status = status

	return service.repository.Update(classroom)

}

func (service *Service) RemoveByID(ID uint) error {

	if ID <= 0 {
		return errors.New(apperror.ClassroomEmptyID)
	}

	classroom, err := service.repository.FindByID(ID)

	if err != nil {
		return err
	}

	return service.repository.Delete(classroom)

}

func generateNewClassroomID() uint {

	const (
		min uint = 100000
		max uint = 999999
	)

	return min + rand.UintN(max-min+1)

}
