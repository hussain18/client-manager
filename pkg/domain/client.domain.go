package domain

import (
	"client-manager/pkg/service"

	"github.com/gin-gonic/gin"
)

func ClientRouters() *gin.Engine {
	r := MainRouter

	r.GET("/clients", service.GetAllClients)

	r.POST("/clients", service.AddClient)

	r.PATCH("/clients/:id", service.UpdateClient)

	r.DELETE("/clients/:id", service.DeleteClient)

	return r
}
