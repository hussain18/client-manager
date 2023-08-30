package service

import (
	"client-manager/pkg/models"
	"client-manager/pkg/repository"
	"client-manager/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	var admin models.Admin

	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	// Get Admin
	dbAdmin, getAdminErr := repository.GetAdmin(admin.Username)
	if getAdminErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to login"})
		log.Println(getAdminErr)
		return
	}

	if dbAdmin == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong username or password"})
		return
	}

	// Compare password
	isPassValid := utils.ComparePassword(admin.Password, dbAdmin.Password)
	if !isPassValid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong username or password"})
		return
	}

	// Generate jwt token
	jwtToken, jwtErr := utils.GenerateToken(admin.Username)
	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to login"})
		log.Println(jwtErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
