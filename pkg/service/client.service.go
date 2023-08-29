package service

import (
	"client-manager/pkg/models"
	"client-manager/pkg/repository"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
