package domain

import (
	"client-manager/pkg/middleware"
	"client-manager/pkg/service"

	"github.com/gin-gonic/gin"
)

func ClientRouters() *gin.Engine {
	r := MainRouter

	r.GET("/clients", middleware.AuthMiddleware(), service.GetAllClients)

	r.POST("/clients", middleware.AuthMiddleware(), service.AddClient)

	r.PATCH("/clients/:id", middleware.AuthMiddleware(), service.UpdateClient)

	r.GET("/clients/:id", middleware.AuthMiddleware(), service.GetClient)

	r.PATCH("/clients/deactivate/:id", middleware.AuthMiddleware(), service.DeReActivateClient)

	r.DELETE("/clients/:id", middleware.AuthMiddleware(), service.DeleteClient)

	return r
}
