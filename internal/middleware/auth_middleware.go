package middleware

import (
	"github/ericoliveiras/basic-crud-go/internal/config"
	"github/ericoliveiras/basic-crud-go/internal/handlers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var jwtSecret = config.LoadAuthConfig().AccessSecret

func AuthMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ERROR::": "Missing Token"})
		ctx.Abort()
		return
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"ERROR::": "Ivalid length or missing prefix bearer"})
		ctx.Abort()
		return
	}

	tokenString := bearerToken[1]
	claims, err := handlers.VerifyJWT(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ERROR::": "Invalid Token"})
		ctx.Abort()
		return
	}

	ctx.Set("claims", claims)

	ctx.Next()
}
