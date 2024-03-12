package professor

import (
	"daltonypng/show-do-humilhao/internal/apperror"
	"daltonypng/show-do-humilhao/internal/entity"
	"errors"
	"regexp"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) Login(email string, password string) error {

	if len(email) <= 0 {
		return errors.New(apperror.ProfessorInvalidEmail)

	} else if len(password) <= 0 {
		return errors.New(apperror.ProfessorInvalidPassword)

	}

	err := service.repository.FindByCredentials(email, password)

	if err != nil {
		return errors.New(apperror.ProfessorUnauthorized)

	}

	return nil

}

func (service *Service) SignIn(professor *entity.Professor) error {

	if !validateEmail(professor.Email) {
		return errors.New(apperror.ProfessorInvalidEmail)

	} else if len(professor.Name) <= 0 {
		return errors.New(apperror.ProfessorEmptyName)

	} else if len(professor.Password) <= 0 {
		return errors.New(apperror.ProfessorInvalidPassword)

	}

	return service.repository.Create(professor)

}

var validEmailRegex regexp.Regexp = *regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)

func validateEmail(email string) bool {
	return validEmailRegex.MatchString(email)
}
