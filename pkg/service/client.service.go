package service

import (
	"client-manager/pkg/models"
	"client-manager/pkg/repository"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Add new client
func AddClient(c *gin.Context) {
	var newClient models.Client

	if err := c.BindJSON(&newClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	// TODO: avoid adding duplicate clients

	if err := repository.AddClient(newClient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add new client"})
		log.Println(err)
		return
	}

	// TODO: add create history for the client

	c.JSON(http.StatusOK, gin.H{"message": "Client added successfully"})
}

// Get all clients
func GetAllClients(c *gin.Context) {
	// TODO: add pagination & search functionality

	clients, err := repository.GetAllClients(*options.Find())
	defer clients.Close(context.TODO())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get clients"})
		log.Println(err)
		return
	}

	var result []bson.M

	if err := clients.All(context.Background(), &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get clients"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateClient(c *gin.Context) {
	id := c.Params.ByName("id")
	var updateBody primitive.D
	var jsonData map[string]interface{}

	if err := c.BindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	for key, value := range jsonData {
		updateBody = append(updateBody, primitive.E{Key: key, Value: value})
	}

	updateResult, err := repository.UpdateClient(id, updateBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update"})
		log.Println(err)
		return
	}

	message := fmt.Sprintf("Matched Items: %d\n Modified Items: %d", updateResult.MatchedCount, updateResult.ModifiedCount)

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func DeleteClient(c *gin.Context) {
	id := c.Params.ByName("id")

	deleteCount, err := repository.DeleteClient(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete"})
		log.Println(err)
		return
	}

	message := fmt.Sprintf("Records deleted: %d", deleteCount)

	c.JSON(http.StatusOK, gin.H{"message": message})
}
