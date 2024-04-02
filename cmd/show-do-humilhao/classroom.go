package main

import (
	"daltonypng/show-do-humilhao/internal/apperror"
	"daltonypng/show-do-humilhao/internal/classroom"
	"daltonypng/show-do-humilhao/internal/covenant"
	"daltonypng/show-do-humilhao/internal/entity"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassroomRouter struct {
	service *classroom.Service
}

func NewClassroomRouter() *ClassroomRouter {

	migration := &entity.Classroom{}

	DSN := os.Getenv("DSN")

	database, err := covenant.NewGorm(&migration, DSN)

	if err != nil {
		panic(err)
	}

	repository := classroom.NewRepository(database)
	service := classroom.NewService(repository)

	return &ClassroomRouter{
		service: service,
	}

}

func (router *ClassroomRouter) postCreateRoom(context *gin.Context) {

	type requestBody struct {
		Name        string `json:"name"`
		ProfessorID uint   `json:"professorId"`
		Status      int    `json:"status"`
	}

	request := &requestBody{}

	err := context.BindJSON(&request)

	if err != nil {
		context.JSON(http.StatusBadRequest, apperror.ClassroomBadRequest)
		return
	}

	classroom := &entity.Classroom{
		Name:        request.Name,
		ProfessorID: request.ProfessorID,
		Status:      request.Status,
	}

	err = router.service.Create(classroom)

	if err != nil {
		status, message := getErrorStatusResponse(err)
		context.String(status, message)
		return
	}

	context.String(http.StatusCreated, strconv.Itoa(int(classroom.ID)))

}

func (router *ClassroomRouter) getClassroom(context *gin.Context) {

	type responseBody struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		ProfessorID uint   `json:"professorId"`
		Status      int    `json:"status"`
	}

	ID, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.String(http.StatusNotFound, "ID de sala inválida.")
		return
	}

	classroom, err := router.service.FindByID(uint(ID))

	if err != nil {
		status, message := getErrorStatusResponse(err)
		context.String(status, message)
		return
	}

	response := &responseBody{
		ID:          classroom.ID,
		Name:        classroom.Name,
		ProfessorID: classroom.ProfessorID,
		Status:      classroom.Status,
	}

	context.JSON(http.StatusOK, response)

}

func (router *ClassroomRouter) putClassroomUpdateStatus(context *gin.Context) {

	type requestBody struct {
		Status int `json:"status"`
	}

	request := &requestBody{}

	err := context.BindJSON(&request)

	if err != nil {
		context.JSON(http.StatusBadRequest, apperror.ClassroomBadRequest)
		return
	}

	ID, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.String(http.StatusNotFound, "ID de sala inválida.")
		return
	}

	err = router.service.UpdateStatusByID(uint(ID), request.Status)

	if err != nil {
		status, message := getErrorStatusResponse(err)
		context.String(status, message)
		return
	}

	context.String(http.StatusOK, "Status da sala atualizado.")

}

func (router *ClassroomRouter) deleteClassroomByID(context *gin.Context) {

	ID, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.String(http.StatusNotFound, "ID de sala inválida.")
		return
	}

	err = router.service.RemoveByID(uint(ID))

	if err != nil {
		status, message := getErrorStatusResponse(err)
		context.String(status, message)
		return
	}

	context.String(http.StatusOK, "Sala removida com sucesso.")

}
