package repository

import (
	"client-manager/pkg/models"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateAdmin() {
	collection := Database.Collection("admin")

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
