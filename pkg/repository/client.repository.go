package repository

import (
	"client-manager/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Add client
func AddClient(clientData models.Client) error {

	collection := Database.Collection("clients")

	_, err := collection.InsertOne(context.Background(), clientData)
	if err != nil {
		return err
	}

	return nil
}

// Update client
// Get clients
func GetAllClients(opts options.FindOptions) (*mongo.Cursor, error) {
	collection := Database.Collection("clients")

	clients, err := collection.Find(context.TODO(), bson.D{{}}, &opts)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

// Search clients
// Get change history
