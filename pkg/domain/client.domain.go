package domain

import (
	"client-manager/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClientRouters() *gin.Engine {
	r := MainRouter

	r.GET("/clients", service.GetAllClients)

	r.POST("/clients", service.AddClient)

	r.PATCH("/clients/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "at PATCH /clients/:id"})
	})

	r.DELETE("/clients/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "at DELETE /clients/:id"})
	})

	return r
}
