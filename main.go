package main

import (
	"client-manager/pkg/domain"
	"client-manager/pkg/repository"
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	domain.RunRouter()
	router := domain.MainRouter

	repository.ConnectDatabase()
	repository.CreateAdmin()

	router.GET("/ping", rootRoute)
	domain.ClientRouters()

	router.Run("localhost:8080")
	defer repository.Client.Disconnect(context.Background())
}

func rootRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "API WORKS!"})
}
