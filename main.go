package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	connectDatabase()

	router.GET("/", rootRoute)

	router.Run("localhost:8080")
}

func connectDatabase() {
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalln("DB_CONN_ERR:", err)
	}

	log.Println("DB connection successful")

	defer client.Disconnect(context.Background())
}

func rootRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "API WORKS!"})
}
