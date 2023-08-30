package middleware

import (
	"client-manager/pkg/repository"
	"client-manager/pkg/utils"
	"log"
	"net/http"
	"strings"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, "Bearer ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid bearer token format"})
			c.Abort()
			return
		}

		tokenString := bearerToken[1]

		token, err := utils.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// If token is valid, set the username in the context for later use
		claims := token.Claims.(jwt.MapClaims)
		username, ok := claims["username"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username in token"})
			c.Abort()
			return
		}

		dbAdmin, getAdminErr := repository.GetAdmin(username)
		if getAdminErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to login"})
			log.Println(getAdminErr)
			return
		}

		if dbAdmin == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Uauthorized Admin"})
			return
		}

		c.Set("admin_username", username)

		c.Next()
	}
}
