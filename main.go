package main

import (
	"client-manager/pkg/domain"
	"client-manager/pkg/repository"
	"context"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
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

	// Define CORS options
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"*"}
	router.Use(cors.New(config))

	repository.ConnectDatabase()
	repository.CreateAdmin()

	router.GET("/ping", rootRoute)
	domain.ClientRouters()
	domain.AdminRouters()

	router.Run("localhost:8080")
	defer repository.Client.Disconnect(context.Background())
}

func rootRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "API WORKS!"})
}
