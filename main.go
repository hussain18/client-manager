package main

import (
	"client-manager/pkg/models"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var database *mongo.Database

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	connectDatabase()
	createAdmin()

	router.GET("/", rootRoute)

	router.Run("localhost:8080")
	defer client.Disconnect(context.Background())
}

func connectDatabase() {
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	clientLocal, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalln("DB_CONN_ERR:", err)
	}

	log.Println("DB connection successful")

	client = clientLocal
	database = client.Database("client-manager")
}

func createAdmin() {
	collection := database.Collection("admin")

	exists, _ := collection.Find(context.TODO(), bson.D{{}})
	defer exists.Close(context.TODO())

	if exists.Next(context.TODO()) {
		return
	}

	admin := models.Admin{
		Username: os.Getenv("ADMIN_USERNAME"),
		Password: os.Getenv("ADMIN_PASSWORD"),
	}

	// Insert the document into the collection
	_, err := collection.InsertOne(context.Background(), admin)
	if err != nil {
		log.Fatal("ADMIN_CREATION_ERROR", err)
	}

	log.Println("Admin created")
}

func rootRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "API WORKS!"})
}
