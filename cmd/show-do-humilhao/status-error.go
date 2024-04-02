package main

import (
	"daltonypng/show-do-humilhao/internal/apperror"
	"net/http"
)

var errorStatus map[string]int = makeErrorStatus()

func makeErrorStatus() map[string]int {

	errors := map[string]int{
		apperror.ProfessorInvalidEmail:     http.StatusBadRequest,
		apperror.ProfessorEmptyName:        http.StatusBadRequest,
		apperror.ProfessorInvalidPassword:  http.StatusBadRequest,
		apperror.ProfessorBadRequest:       http.StatusBadRequest,
		apperror.ProfessorDuplicated:       http.StatusConflict,
		apperror.ProfessorUnauthorized:     http.StatusUnauthorized,
		apperror.ClassroomBadRequest:       http.StatusBadRequest,
		apperror.ClassroomEmptyName:        http.StatusBadRequest,
		apperror.ClassroomInvalidProfessor: http.StatusBadRequest,
		apperror.ClassroomEmptyID:          http.StatusBadRequest,
		apperror.ClassroomInvalidStatus:    http.StatusBadRequest,
		apperror.ClassroomNotFound:         http.StatusNotFound,
	}

	return errors
}

func getErrorStatusResponse(err error) (int, string) {

	message := err.Error()
	status := errorStatus[message]

	if status > 0 {
		return status, message
	}

	return http.StatusInternalServerError, "Erro não tratado: " + message
}
