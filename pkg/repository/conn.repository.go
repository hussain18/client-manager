package repository

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func ConnectDatabase() {
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	clientLocal, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalln("DB_CONN_ERR:", err)
	}

	log.Println("DB connection successful")

	Client = clientLocal
	Database = clientLocal.Database(os.Getenv("DB_NAME"))
}
