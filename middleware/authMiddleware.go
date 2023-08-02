package middleware

import (
	"github/ericoliveiras/basic-crud-go/handlers"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var err = godotenv.Load()
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ERROR::": "Missing Token"})
		c.Abort()
		return
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"ERROR::": "Invalid token"})
		c.Abort()
		return
	}

	tokenString := bearerToken[1]
	claims, err := handlers.VerifyJWT(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ERROR::": "Invalid Token"})
		c.Abort()
		return
	}

	c.Set("claims", claims)

	c.Next()
}
