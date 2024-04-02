package classroom_test

import (
	"daltonypng/show-do-humilhao/internal/apperror"
	"daltonypng/show-do-humilhao/internal/classroom"
	"daltonypng/show-do-humilhao/internal/covenant"
	"daltonypng/show-do-humilhao/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

var repository *classroom.Repository = classroom.NewRepository(covenant.NewMocker())

func TestCreateRoomEmptyProfessor(t *testing.T) {

	service := classroom.NewService(repository)

	classroom := &entity.Classroom{
		Name: "Test Classroom",
	}

	err := service.Create(classroom)

	assert.NotNil(t, err)
	assert.EqualError(t, err, apperror.ClassroomInvalidProfessor)
}

func TestCreateRoom(t *testing.T) {
	var min uint = 100000
	var max uint = 999999

	service := classroom.NewService(repository)

	classroom := &entity.Classroom{
		Name:        "Test Classroom",
		ProfessorID: 1,
	}

	err := service.Create(classroom)

	assert.Nil(t, err)
	assert.Greater(t, classroom.ID, min)
	assert.Less(t, classroom.ID, max)
}

func TestRemoveRoomByID(t *testing.T) {

	service := classroom.NewService(repository)

	classroom := &entity.Classroom{
		Name:        "Test Classroom",
		ProfessorID: 1,
	}

	err := service.Create(classroom)

	assert.Nil(t, err)
	assert.NotNil(t, classroom)

	err = service.RemoveByID(classroom.ID)

	assert.Nil(t, err)

}

func TestFindByID(t *testing.T) {

	service := classroom.NewService(repository)

	classroom, err := service.FindByID(100000)

	assert.Nil(t, err)

	assert.NotNil(t, classroom)
}

func TestUpdateStatusByID(t *testing.T) {

	service := classroom.NewService(repository)

	err := service.UpdateStatusByID(100000, 1)

	assert.Nil(t, err)

	err = service.UpdateStatusByID(100000, 999)

	assert.NotNil(t, err)
	assert.EqualError(t, err, apperror.ClassroomInvalidStatus)

}
