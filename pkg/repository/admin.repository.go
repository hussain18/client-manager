package repository

import (
	"client-manager/pkg/models"
	"client-manager/pkg/utils"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAdmin() {
	collection := Database.Collection("admin")

	exists, _ := collection.Find(context.TODO(), bson.D{{}})
	defer exists.Close(context.TODO())

	if exists.Next(context.TODO()) {
		return
	}

	encryptedPassword, encErr := utils.EncryptPassword(os.Getenv("ADMIN_PASSWORD"))
	if encErr != nil {
		log.Fatal("ADMIN_CREATION_ERROR", encErr)
	}

	admin := models.Admin{
		Username: os.Getenv("ADMIN_USERNAME"),
		Password: encryptedPassword,
	}

	// Insert the document into the collection
	_, err := collection.InsertOne(context.Background(), admin)
	if err != nil {
		log.Fatal("ADMIN_CREATION_ERROR", err)
	}

	log.Println("Admin created")
}

func GetAdmin(username string) (*models.Admin, error) {
	collection := Database.Collection("admin")
	var result models.Admin

	if err := collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&result); err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}
