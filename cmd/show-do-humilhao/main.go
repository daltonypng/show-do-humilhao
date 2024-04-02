package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	server := gin.Default()
	professorRouter := NewProfessorRouter()
	classroomRouter := NewClassroomRouter()

	server.POST("/v1/sign-in", professorRouter.postSignIn)
	server.POST("/v1/login", professorRouter.postLogin)

	server.POST("/v1/classroom", classroomRouter.postCreateRoom)
	server.GET("/v1/classroom/:id", classroomRouter.getClassroom)
	server.PUT("/v1/classroom/:id", classroomRouter.putClassroomUpdateStatus)
	server.DELETE("/v1/classroom/:id", classroomRouter.deleteClassroomByID)

	server.Run()
}
