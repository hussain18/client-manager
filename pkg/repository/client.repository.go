package repository

import (
	"client-manager/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func UpdateClient(id string, updateBody primitive.D) (*mongo.UpdateResult, error) {
	collection := Database.Collection("clients")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.D{{Key: "$set", Value: updateBody}}

	result, updateError := collection.UpdateByID(context.TODO(), objectID, update)

	if updateError != nil {
		return nil, updateError
	}

	return result, nil
}

// Get clients
func GetAllClients(opts options.FindOptions) (*mongo.Cursor, error) {
	collection := Database.Collection("clients")

	clients, err := collection.Find(context.TODO(), bson.D{{}}, &opts)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

// Get one Client
func GetOneClient(filter primitive.D, opts options.FindOneOptions) (*models.Client, error) {
	collection := Database.Collection("clients")
	var result models.Client

	if err := collection.FindOne(context.TODO(), filter, &opts).Decode(&result); err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

// Search clients
// Get change history
// Delete client
func DeleteClient(id string) (int, error) {
	collection := Database.Collection("clients")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := collection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}})

	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}
