package professor_test

import (
	"daltonypng/show-do-humilhao/internal/apperror"
	"daltonypng/show-do-humilhao/internal/covenant"
	"daltonypng/show-do-humilhao/internal/entity"
	"daltonypng/show-do-humilhao/internal/professor"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	mockName     string = "John Doe"
	mockEmail    string = "johndoe@mail.com"
	mockPassword string = "abc123"
)

var repository *professor.Repository = professor.NewRepository(covenant.NewMocker())
var errorRepository *professor.Repository = professor.NewRepository(covenant.NewMockerError())

func TestLoginEmptyEmail(t *testing.T) {

	service := professor.NewService(repository)

	err := service.Login("", "")

	assert.NotNil(t, err)
	assert.EqualError(t, err, apperror.ProfessorInvalidEmail)

}

func TestLoginEmptyPassword(t *testing.T) {

	service := professor.NewService(repository)

	err := service.Login(mockPassword, "")

	assert.NotNil(t, err)
	assert.EqualError(t, err, apperror.ProfessorInvalidPassword)

}

func TestLoginUnauthorized(t *testing.T) {

	service := professor.NewService(errorRepository)

	err := service.Login(mockPassword, mockPassword)

	assert.NotNil(t, err)
	assert.EqualError(t, err, apperror.ProfessorUnauthorized)

}

func TestLogin(t *testing.T) {

	service := professor.NewService(repository)

	err := service.Login(mockPassword, mockPassword)

	assert.Nil(t, err)

}

func TestSignInEmptyEmail(t *testing.T) {

	service := professor.NewService(repository)

	professor := &entity.Professor{
		Email: "",
	}

	err := service.SignIn(professor)

	assert.NotNil(t, err)
	assert.EqualError(t, err, apperror.ProfessorInvalidEmail)

}

func TestSignInvalidEmail(t *testing.T) {

	service := professor.NewService(repository)

	professor := &entity.Professor{
		Email: "totaly_not_a_mail",
	}

	err := service.SignIn(professor)

	assert.NotNil(t, err)
	assert.EqualError(t, err, apperror.ProfessorInvalidEmail)

}

func TestSignInEmptyName(t *testing.T) {

	service := professor.NewService(repository)

	professor := &entity.Professor{
		Email: mockEmail,
		Name:  "",
	}

	err := service.SignIn(professor)

	assert.NotNil(t, err)
	assert.EqualError(t, err, apperror.ProfessorEmptyName)

}

func TestSignInInvalidPassword(t *testing.T) {

	service := professor.NewService(repository)

	professor := &entity.Professor{
		Email:    mockEmail,
		Name:     mockName,
		Password: "",
	}

	err := service.SignIn(professor)

	assert.NotNil(t, err)
	assert.EqualError(t, err, apperror.ProfessorInvalidPassword)

}

func TestSignIn(t *testing.T) {

	service := professor.NewService(repository)

	professor := &entity.Professor{
		Email:    mockEmail,
		Name:     mockName,
		Password: mockPassword,
	}

	err := service.SignIn(professor)

	assert.Nil(t, err)

}
