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

	server.POST("/v1/sign-in", professorRouter.postSignIn)

	server.Run()
}
