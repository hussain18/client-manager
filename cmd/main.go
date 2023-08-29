package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", rootRoute)

	router.Run("localhost:8080")
}

func rootRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "API WORKS!"})
}
