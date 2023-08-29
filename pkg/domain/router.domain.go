package domain

import "github.com/gin-gonic/gin"

var MainRouter *gin.Engine

func RunRouter() {
	MainRouter = gin.Default()
}
