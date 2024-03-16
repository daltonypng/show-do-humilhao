package main

import (
	"daltonypng/show-do-humilhao/internal/apperror"
	"daltonypng/show-do-humilhao/internal/covenant"
	"daltonypng/show-do-humilhao/internal/entity"
	"daltonypng/show-do-humilhao/internal/professor"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ProfessorRouter struct {
	service *professor.Service
}

func NewProfessorRouter() *ProfessorRouter {

	migration := &entity.Professor{}

	DSN := os.Getenv("DSN")

	database, err := covenant.NewGorm(&migration, DSN)

	if err != nil {
		panic(err)
	}

	repository := professor.NewRepository(database)
	service := professor.NewService(repository)

	return &ProfessorRouter{
		service: service,
	}

}

type signInRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (router *ProfessorRouter) postSignIn(context *gin.Context) {

	requestBody := &signInRequestBody{}

	err := context.BindJSON(&requestBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, apperror.ProfessorBadRequest)
		return
	}

	professor := &entity.Professor{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	err = router.service.SignIn(professor)

	if err != nil {
		context.JSON(professorAppErrorStatusCode(err), err)
		return
	}

	context.JSON(http.StatusCreated, "Professor cadastrado com sucesso.")

}

func professorAppErrorStatusCode(err error) int {

	message := err.Error()

	switch message {
	case apperror.ProfessorInvalidEmail:
	case apperror.ProfessorEmptyName:
	case apperror.ProfessorInvalidPassword:
		return http.StatusBadRequest

	}

	return http.StatusInternalServerError
}
