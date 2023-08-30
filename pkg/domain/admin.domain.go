package domain

import (
	"client-manager/pkg/service"

	"github.com/gin-gonic/gin"
)

func AdminRouters() *gin.Engine {
	r := MainRouter

	r.POST("/admin/login", service.AdminLogin)

	return r
}
