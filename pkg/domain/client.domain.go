package domain

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClientRouters() *gin.Engine {
	r := MainRouter

	r.GET("/clients", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "at GET /clients"})
	})

	r.POST("/clients", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "at POST /clients"})
	})

	r.PATCH("/clients/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "at PATCH /clients/:id"})
	})

	r.DELETE("/clients/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "at DELETE /clients/:id"})
	})

	return r
}
