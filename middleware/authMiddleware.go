package middleware

import (
	"github/ericoliveiras/basic-crud-go/handlers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var err = godotenv.Load()
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(c *gin.Context) {
	tokenString, err := handlers.ExtractTokenFromRequest(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ERROR::": "Invalid Token"})
		c.Abort()
		return
	}

	claims, err := handlers.VerifyJWT(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ERROR::": "Invalid Token"})
		c.Abort()
		return
	}

	c.Set("claims", claims)

	c.Next()
}
